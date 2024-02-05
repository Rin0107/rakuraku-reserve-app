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
//　初期パスワードはハッシュ化したpasswordとする
func CreateUsers(name,email,role string,password string)(){
	user:=User{}
	user.Name=name
	user.Email=email
	user.Role=role
	user.Password=password
	
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

// emailからユーザー情報を取得するためのメソッドt
func GetUserPasswordByEmail(email string) (password string,role string){
	user:=User{}
	result:=Db.Select("Password","Role").Where("email=?",email).Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user.Password,user.Role
}

func SaveTokenToUser(email string,token string){
	result:=Db.Table("users").Where("email = ?", email).Updates(User{PasswordResetToken: token,UpdatedAt: time.Now()})
	if result.Error != nil {
		panic(result.Error)
	}
}