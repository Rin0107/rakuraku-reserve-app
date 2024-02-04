package controller

import (
	"app/service"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
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

var (
	// セッション情報を保存するためのマップ
	sessions = make(map[string]bool)
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
		err:=service.Login(loginInformation.Email,loginInformation.Password)
		if err != nil {
			// 存在しないemailまたは適切なpasswordが入力されていない場合403を返す
			errorMessage := ResponseMessage{Message: err.Error()}
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
			sessions[sessionID] = true
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
