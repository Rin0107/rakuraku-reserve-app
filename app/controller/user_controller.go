package controller

import (
	"app/service"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// ログイン情報
type LoginInformation struct {
	Email string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}

// レスポンスメッセージ格納用
type ResponseMessage struct {
	Message string `json:"message"`
}

// パスワードリセットメール送信用のユーザー情報
type UserInformationForResetPassword struct{
	Email string `json:"email" validate:"email,existing_email_validation"`
}
// セッション情報に格納するユーザー情報
type SessionUserInformation struct {
	UserId int
	Email string
	Role string
}

// パスワードリセット処理のユーザー情報
type PasswordResetInformation struct{
	PasswordToken string `json:"password_token" validate:"required"`
	Password string `json:"password" validate:"password_confirmation_validation"`
	ConfirmationPassword string `json:"confirmation_password" validate:"required"`
}

// 変更用のユーザー情報
type UserInformationToUpdate struct{
	Name  string `json:"name" validate:"required,gte=0,lte=100"`
	Email string `json:"email" validate:"email"`
	UserIcon string `json:"user_icon"`
}
var (
	// セッション情報を保存するためのマップ
	sessions = make(map[string]SessionUserInformation)
	// セッション情報へのアクセスを同期するためのミューテックス
	sessionMutex = &sync.Mutex{}
)

// ユーザー一覧を取得するためのメソッド
// ユーザーが存在する場合はユーザー情報をリストにして200で返す
// ユーザーが存在しない場合は空のリストを404で返す
func GetUsers(c *gin.Context){
	users := service.GetUsers()
	switch len(users){
	case 0:
		c.IndentedJSON(404, users)
	default:
		c.IndentedJSON(200, users)
	}
}

// ログイン処理のためのメソッド
//　ログイン情報を受け取りログイン処理を実施する
// バリデーションとしてemailはemail形式,passwordは入力必須とした
// ログイン成功時200を返す（認証情報については後ほど追記する）
// 入力形式が不適切の場合400を返す
// 入力値が不適切であれば403を返す
func Login(c *gin.Context){
	var loginInformation LoginInformation
	if err := c.ShouldBindJSON(&loginInformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 以下バリデーション実装
	validate = validator.New()
	validateErr := validate.Struct(loginInformation)
	if validateErr!=nil{
		for _, err := range validateErr.(validator.ValidationErrors) {
			fmt.Println("Namespace =",err.Namespace())
			fmt.Println("Tag =",err.Tag())
			fmt.Println("Type =",err.Type())
			fmt.Println("Value =",err.Value())
			fmt.Println("Param =",err.Param())
		}
		errorMessage :=ResponseMessage{Message: "ログインが失敗しました"}
		c.IndentedJSON(400, errorMessage)
	}else{
		//ログイン処理を呼び出す
		userId,role,err:=service.Login(loginInformation.Email,loginInformation.Password)
		if err != nil {
			// 存在しないemailまたは適切なpasswordが入力されていない場合403を返す
			errorMessage := ResponseMessage{Message: "ユーザ ID 又はパスワードが不正です"}
			c.JSON(403,errorMessage)
		}else{
			// 認可情報としてセッションIDを作成する
			sessionID,err:=generateSessionID()
			if err != nil {
				errorMessage := ResponseMessage{Message: err.Error()}
				c.JSON(500,errorMessage)
			}

			// セッションIDを保存する
			sessionMutex.Lock()
			sessionUserInformation:=SessionUserInformation{UserId: userId,Email: loginInformation.Email,Role: role}
			sessions[sessionID] = sessionUserInformation
			sessionMutex.Unlock()

			// Cookieをセットする
			cookie := &http.Cookie{
				Name: "session_id",
				Value: sessionID,
				Path: "/",
				MaxAge: 3600,
			}
			http.SetCookie(c.Writer, cookie)

			// 認証処理が成功の場合200を返す
			message := ResponseMessage{Message: "ログインに成功しました"}
			c.JSON(200, message)
		}
	}
}

// ランダムにセッションIDを作成するためのメソッド
func generateSessionID() (string, error) {
   	// 16バイトのランダムなバイト列を生成
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// base64エンコードしてセッションIDとして使用
	sessionID := base64.URLEncoding.EncodeToString(randomBytes)
	return sessionID, nil
}

// ログアウトするためのメソッド
// サーバー、クライアント側それぞれのセッションID情報を削除する
func Logout(c *gin.Context) {
	// CookieからセッションIDを取得し、サーバーサイドでセッションを削除
	cookie, err := c.Cookie("session_id")
	if err == nil {
		sessionMutex.Lock()
		delete(sessions, cookie)
		sessionMutex.Unlock()
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Cookieを削除
	})

	message := ResponseMessage{Message: "適切にログアウトされました"}
	c.JSON(200,message)
}

// パスワードを変更するためにメールを送信するためのメソッド
// メールアドレスがリクエストボディに必要
// 本メソッドでemail形式、存在するemailか確認する
// 成功時はメール送信後、200で返す
// メール送信失敗時、500で返す
// バリデーション時、400で返す
func SendEmailToChangePassword(c *gin.Context){
	var userInformationForResetPassword UserInformationForResetPassword
	if err := c.ShouldBindJSON(&userInformationForResetPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 以下バリデーション実装
	validate = validator.New()
	validate.RegisterValidation("existing_email_validation",existingEmailValidation)
	validateErr := validate.Struct(userInformationForResetPassword)
	if validateErr!=nil{
		for _, err := range validateErr.(validator.ValidationErrors) {
			fmt.Println("Namespace =",err.Namespace())
			fmt.Println("Tag =",err.Tag())
			fmt.Println("Type =",err.Type())
			fmt.Println("Value =",err.Value())
			fmt.Println("Param =",err.Param())
		} 
		var errorMessage Error
		errorMessage.Message="不正なメールアドレスです"
		c.IndentedJSON(400, errorMessage)
	}else{
		err:=service.SendEmailToChangePassword(userInformationForResetPassword.Email)
		if err != nil {
			var errorMessage Error
			errorMessage.Message="メール送信に失敗しました"
			c.JSON(500,errorMessage)
		}else{
			// メールアドレスが問題ない場合、メール送信処理を呼び出す
			c.IndentedJSON(200, userInformationForResetPassword)
		}
	}
}

// パスワード再設定のためのメソッド
// メールに添付されているトークンが必要
// パスワードと確認用パスワードが同値である必要がある
func ResetPassword(c *gin.Context){
	var passwordResetInformation PasswordResetInformation
	if err := c.ShouldBindJSON(&passwordResetInformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 以下バリデーション実装
	validate = validator.New()
	validate.RegisterValidation("password_confirmation_validation",passwordConfirmationValidation)
	validateErr := validate.Struct(passwordResetInformation)
	if validateErr!=nil{
		for _, err := range validateErr.(validator.ValidationErrors) {
			fmt.Println("Namespace =",err.Namespace())
			fmt.Println("Tag =",err.Tag())
			fmt.Println("Type =",err.Type())
			fmt.Println("Value =",err.Value())
			fmt.Println("Param =",err.Param())
		} 
		var errorMessage Error
		errorMessage.Message="不正な入力値があります"
		c.IndentedJSON(400, errorMessage)
	}else{
		err:=service.ResetPassword(passwordResetInformation.PasswordToken,passwordResetInformation.Password)
		if err != nil {
			var errorMessage Error
			errorMessage.Message="パスワード再設定が失敗しました"
			c.JSON(500,errorMessage)
		}else{
			// パスワードが再設定されたらログイン状態にする
			message := ResponseMessage{Message: "パスワードが再設定されました"}
			c.IndentedJSON(200, message)
		}
	}
}

// ユーザー詳細を取得するためのメソッド
func GetUserDetail(c *gin.Context){
	// URLからuserIdを取得する
	userId,_ := strconv.Atoi(c.Param("userId"))
	//　パラメータがない場合（/api/user）の場合、自身のユーザー詳細を返す
	if userId == 0 {
		// sessionIdからuserIdを取得する
		userId=GetUserInformationBySessionId(c).UserId
	}
	// userIdからユーザー詳細を取得する
	userDetail,err:=service.GetUserDetail(userId)
	if err != nil {
		fmt.Println(err)
		errorMessage := ResponseMessage{Message: "ユーザー情報取得に失敗しました"}
		c.JSON(500,errorMessage)
		return
	}
	c.JSON(200,userDetail)
}

// ユーザー情報を変更するためのメソッド
// 管理者権限の場合、すべてのユーザーのすべてのユーザー情報を変更できる
// ユーザー権限の場合、自分のユーザー情報（role以外）を変更できる
func UpdateUserInformation(c *gin.Context){
	// URLからuserIdを取得する
	userId,_ := strconv.Atoi(c.Param("userId"))
	
	// リクエストボディを取得する
	var userInformationToUpdate UserInformationToUpdate
	if err := c.ShouldBindJSON(&userInformationToUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 以下バリデーション実装
	validate = validator.New()
	validateErr := validate.Struct(userInformationToUpdate)
	if validateErr!=nil{
		for _, err := range validateErr.(validator.ValidationErrors) {
			fmt.Println("Namespace =",err.Namespace())
			fmt.Println("Tag =",err.Tag())
			fmt.Println("Type =",err.Type())
			fmt.Println("Value =",err.Value())
			fmt.Println("Param =",err.Param())
		} 
		var errorMessage Error
		errorMessage.Message="不正な入力値があります"
		c.IndentedJSON(400, errorMessage)
		return
	}
	
	// session情報の権限情報がuserの場合、自分のユーザー情報のみを変更できる
	if GetUserInformationBySessionId(c).Role=="user" {
		if GetUserInformationBySessionId(c).UserId!=userId {
			var errorMessage Error
			errorMessage.Message="不正な入力値があります"
			c.IndentedJSON(400, errorMessage)
			return
		}
	} 

	// 自分のメール以外が指定されたとき、重複したメールがあるか確認する
	userDetail,getUserDetailErr:=service.GetUserDetail(userId)
	if getUserDetailErr != nil {
		fmt.Println(getUserDetailErr)
		errorMessage := ResponseMessage{Message: "ユーザー情報取得に失敗しました"}
		c.JSON(500,errorMessage)
		return
	}
	if GetUserInformationBySessionId(c).Email!=userDetail.Email {
		if !service.IsEmail(userInformationToUpdate.Email) {
			var errorMessage Error
			errorMessage.Message="ユーザー情報の変更に失敗しました"
			c.IndentedJSON(400, errorMessage)
			return
		}
	}
	
	// userIdを指定してユーザー情報を変更する
	err := service.UpdateUserInformation(userId,userInformationToUpdate.Name,userInformationToUpdate.Email,userInformationToUpdate.UserIcon)
	if err != nil {
		var errorMessage Error
		errorMessage.Message="ユーザー情報の変更に失敗しました"
		c.JSON(500,errorMessage)
		return
	}
	message := ResponseMessage{Message: "ユーザー情報が変更されました"}
	c.IndentedJSON(200, message)
}

//存在するメールアドレスがあるか確認するカスタムバリデーション実装
func existingEmailValidation(fl validator.FieldLevel) bool{
	email := fl.Field().String()
	return !service.IsEmail(email)
}

// パスワードと確認用パスワードが一致するか確認するカスタムバリデーション
func passwordConfirmationValidation(fl validator.FieldLevel) bool {
    // パスワードと確認用パスワードを取得
    password := fl.Field().String()
	confirmPasswordField := fl.Parent().FieldByName("ConfirmationPassword")

    // 確認用パスワードが存在しない場合は一致しないとみなす
    if !confirmPasswordField.IsValid() {
        return false
    }

    confirmPassword := confirmPasswordField.String()

    // パスワードと確認用パスワードが一致するか確認
    return password == confirmPassword
}

// セッション情報からユーザー情報を取得するための汎用的メソッド
func GetUserInformationBySessionId(c *gin.Context) SessionUserInformation{
	sessionId, err := c.Cookie("session_id")
		if err != nil {
			fmt.Println(err)
			errorMessage := ResponseMessage{Message: "ユーザー情報取得に失敗しました"}
			c.JSON(500,errorMessage)
			return SessionUserInformation{}
		}
		// ログインユーザーのユーザーIDを取得
		sessionMutex.Lock()
		sessionUserInformation := sessions[sessionId]
		sessionMutex.Unlock()

		return sessionUserInformation
}