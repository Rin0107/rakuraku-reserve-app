package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Password string
}

func GetAll() (users []User) {
	result := Db.Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (u *User) Create() {
	result := Db.Create(u)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

