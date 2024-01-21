package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	UserId int
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

func GetAll() (users []User) {
	result := Db.Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

// func (u *User) Create() {
// 	result := Db.Create(u)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }

func GetUsers()(users []User){
	result := Db.Where("deleted_at IS NULL").Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}