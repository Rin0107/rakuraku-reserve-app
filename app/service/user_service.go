package service

import (
	"app/model"
	"errors"

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
		return 0,"",errors.New("存在しないメールアドレスです")
	}
	err := CompareHashAndPassword(userPassword, password)
	if err != nil {
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