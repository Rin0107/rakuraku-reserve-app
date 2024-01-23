package service

import (
	"app/model"
)

type User struct {
	Name  string
	Email string
	Role  string
}

func GetUsers() []model.User{
	users := model.GetUsers()
	return users
}

func CreateUsers(name,email,role string){
	model.CreateUsers(name,email,role)
}