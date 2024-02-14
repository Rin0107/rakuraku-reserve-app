package service

import (
	"fmt"
	"mime"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type MailInformation struct{
	Email string `validate:"email"` // メールアドレス
	Subject string `validate:"required"` // 件名
	Text string `validate:"required"` // 本文
}

// メール送信のためのメソッド
// MailInformationを引数に取り、メールを作成する
func SendMail(mailInformation MailInformation)error{
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("読み込み出来ませんでした: %v", envErr)
	} 
	//メール送信処理
	// 送信元メールアドレス
	from := os.Getenv("MAIL_ADDRESS")
	// メールタイトルのエンコード
	subject := mime.QEncoding.Encode("utf-8", mailInformation.Subject)
	// SMTPサーバーのアドレスとポート（Gmailの場合）
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	// 送信元メールアドレスのユーザー名とパスワード（Gmailの場合はアプリパスワードを推奨）
	smtpUsername := os.Getenv("MAIL_ADDRESS")
	smtpPassword := os.Getenv("GMAIL_PASSWORD")

	// メール内容詳細
	// メールの構築
	message := []byte("To: " + mailInformation.Email + "\r\n" +
		"Subject: "+subject+"\r\n\r\n" +
		"\r\n" +
		mailInformation.Text+"\n\n\n"+
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
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, from, []string{mailInformation.Email}, []byte(message))
	if err != nil {
		return err
	}

	fmt.Println("Mail sent successfully.")
	return nil
}