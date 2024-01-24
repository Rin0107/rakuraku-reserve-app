package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

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

func GetUsers()(users []User){
	result := Db.Where("deleted_at IS NULL").Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

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

func IsEmail(email string) (users []User){
	result := Db.Where("email=?",email).Find(&users)
	fmt.Println(email)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}