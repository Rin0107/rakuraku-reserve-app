package service

import (
	"app/model"
)

//ユーザー一覧を取得するためのメソッド
// ユーザー情報を返す
func GetUsers() []model.User{
	users := model.GetUsers()
	return users
}

//ユーザー登録するためのメソッド
func CreateUsers(name,email,role string){
	model.CreateUsers(name,email,role)
}

//メール重複を確認するためのメソッド
func IsEmail(email string) bool{
	user := model.IsEmail(email)
	return len(user) == 0
}