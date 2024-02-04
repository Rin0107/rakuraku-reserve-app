package service

import (
	"app/model"
	"errors"
	"fmt"
	"net/smtp"

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
	//初期パスワードをハッシュ化したpasswordとして生成
	password:="password"
	encryptPw,err:=PasswordEncrypt(password)
	if err != nil {
		fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
		return
	}
	model.CreateUsers(name,email,role,encryptPw)
}

// ログイン処理のためのメソッド
func Login(email string,password string)(role string,error error){
	//　入力されたログイン情報から認証処理を実施する
	// emailからユーザー情報（password）を取得する
	userPassword,role :=model.GetUserPasswordByEmail(email)
	if userPassword == "" {
		return "",errors.New("存在しないメールアドレスです")
	}
	err := CompareHashAndPassword(userPassword, password)
	if err != nil {
		return "",errors.New("パスワードが一致しませんでした")
	}
	return role,nil
}

//メール重複を確認するためのメソッド
func IsEmail(email string) bool{
	user := model.IsEmail(email)
	return len(user) == 0
}

// 暗号(Hash)化するためのメソッド
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 暗号(Hash)と入力された平パスワードの比較
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

//　パスワードをリセットするためのメールを送信するためのメソッド
// メール送信が失敗した場合、エラーを返す
func SendEmailToChangePassword(email string) error{
	//メール送信処理
	// 送信元メールアドレス
	from := "kazuo.ikeda0716@gmail.com"
	// 送信先メールアドレス
	to := email
	// SMTPサーバーのアドレスとポート（Gmailの場合）
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	// 送信元メールアドレスのユーザー名とパスワード（Gmailの場合はアプリパスワードを推奨）
	smtpUsername := "kazuo.ikeda0716@gmail.com"
	smtpPassword := "mekabu1411"

	// メールの構築
	message := []byte("To: " + to + "\r\n" +
		"Subject: Test Subject\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	// SMTPサーバーに接続
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, from, []string{to}, message)
	if err != nil {
		return err
	}

	fmt.Println("Mail sent successfully.")
	return nil
}