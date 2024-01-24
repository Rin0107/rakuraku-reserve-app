package service

import (
	"app/model"
	"fmt"
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

func IsEmail(email string) bool{
	user := model.IsEmail(email)
	fmt.Println(user)
	return len(user) == 0
}