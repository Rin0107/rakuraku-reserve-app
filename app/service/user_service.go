package service

import (
	"app/model"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"mime"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"

	"golang.org/x/crypto/bcrypt"
)

//ユーザー一覧を取得するためのメソッド
// ユーザー情報を返す
func GetUsers() []model.User{
	users := model.GetUsers()
	return users
}

//ユーザー登録するためのメソッド
func CreateUsers(name,email,role string){
	// 初期パスワードをハッシュ化したpasswordとして生成
	// パスワードのハッシュ化はmodelに実装している
	password:="password"
	model.CreateUsers(name,email,role,password)
}

// ログイン処理のためのメソッド
func Login(email string,password string)(userId int,role string,error error){
	//　入力されたログイン情報から認証処理を実施する
	// emailからユーザー情報（password）を取得する
	userPassword,userId,role :=model.GetUserPasswordByEmail(email)
	if userPassword == "" {
		fmt.Println("認証時にエラーが発生しました: ",userPassword)
		return 0,"",errors.New("存在しないメールアドレスです")
	}
	err := CompareHashAndPassword(userPassword, password)
	if err != nil {
		fmt.Println("認証時にエラーが発生しました: ",err)
		return 0,"",errors.New("パスワードが一致しませんでした")
	}
	return userId,role,nil
}

//メール重複を確認するためのメソッド
func IsEmail(email string) bool{
	user := model.IsEmail(email)
	return len(user) == 0
}


// 暗号(Hash)と入力された平パスワードの比較
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

//　パスワードをリセットするためのメールを送信するためのメソッド
// メール送信が失敗した場合、エラーを返す
func SendEmailToChangePassword(email string) error{
	
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("読み込み出来ませんでした: %v", envErr)
	} 
	//メール送信処理
	// 送信元メールアドレス
	from := os.Getenv("MAIL_ADDRESS")
	// 送信先メールアドレス
	to := email
	// メールタイトル
	subject:="【楽々予約】パスワードの再設定について"
	subject = mime.QEncoding.Encode("utf-8", subject)
	// SMTPサーバーのアドレスとポート（Gmailの場合）
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	// 送信元メールアドレスのユーザー名とパスワード（Gmailの場合はアプリパスワードを推奨）
	smtpUsername := os.Getenv("MAIL_ADDRESS")
	smtpPassword := os.Getenv("GMAIL_PASSWORD")

	// メール内容詳細
	// トークン作成処理
	// 作成後、ユーザーテーブルに挿入
	token, tokenErr := generateRandomToken(10)
	if tokenErr != nil {
		fmt.Println("Error generating token:", tokenErr)
		return tokenErr
	}
	model.SaveTokenToUser(email,token)

	// メールの構築
	message := []byte("To: " + to + "\r\n" +
		"Subject: "+subject+"\r\n\r\n" +
		"\r\n" +
		email+"宛にパスワードの再設定がリクエストされました。\r\n"+
		"以下のトークンを使用して再設定が可能です。\n" + 
		token + "\n\n"+
		"このメールに心当たりが無い場合は無視してください。\n"+
		"上記トークンを通して再設定しない限り、パスワードは変更されません。"+"\n\n\n"+
  		//フッター
		"------------------------------------------------------\r\n" + 
		"株式会社ラクスパートナーズ\r\n" + 
		"楽々予約システム\r\n" + 
		"〒160-0022\r\n" + 
		"東京都新宿区新宿4-3-25\r\n"+
		"TOKYU REIT新宿ビル8F\r\n" + 
		"TEL：03-6675-3638\r\n"+ 
		"FAX： 0120-82-5349\r\n" + 
		"E-MAIL: "+os.Getenv("MAIL_ADDRESS")+"\r\n" + 
		"------------------------------------------------------")

	// SMTPサーバーに接続
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, from, []string{to}, message)
	if err != nil {
		return err
	}

	fmt.Println("Mail sent successfully.")
	return nil
}

// ランダムなトークンを作成する処理
// 文字数を引数にとり、適切なトークンを作成する
func generateRandomToken(length int) (string, error) {
	// 生成するランダムなバイト列の長さ
	randomBytes := make([]byte, length)

	// crypto/rand パッケージを使用してランダムなバイト列を生成
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// base64エンコードしてトークンとして使用
	token := base64.RawURLEncoding.EncodeToString(randomBytes)
	return token, nil
}

// パスワードを再設定するためのメソッド
func ResetPassword(passwordToken string,password string) error{
	// トークンを持ったユーザーのユーザーIDを取得する
	userId,err:=model.GetUserIdForPasswordToken(passwordToken)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// ユーザーIDを使ってパスワードを変更する
	model.ResetPassword(userId,password)

	// ユーザーIDを使ってトークン情報を削除する
	model.DeleteResetPasswordToken(userId)

	return nil
}