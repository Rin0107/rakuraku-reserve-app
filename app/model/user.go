package model

import (
	"time"

	"gorm.io/gorm"
)

// usersテーブル情報
// userIdをprimaryKeyに指定している
type User struct {
	UserId int `gorm:"primaryKey;autoIncrement"`
	Name string
	Email string
	Password string
	UserIcon string
	Role string
	PasswordResetToken string
	CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt
}

// DBからユーザー一覧を取得するためのメソッド
// 削除されていないユーザーを取得（deleted_at IS NULL）
func GetUsers()(users []User){
	result := Db.Where("deleted_at IS NULL").Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

// ユーザーを登録するためのメソッド
// パスワードをデフォルトでpasswordに設定（仕様にて変更あり）
func CreateUsers(name,email,role string)(){
	user:=User{}
	user.Name=name
	user.Email=email
	user.Role=role
	user.Password="password"
	
	result:=Db.Select("Name","Email","Role","Password").Create(&user)
	if result.Error != nil {
	panic(result.Error)
	}
	return
}

// メール情報からユーザー情報を取得するためのメソッド
func IsEmail(email string) (users []User){
	result := Db.Where("email=?",email).Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}