package service

import(
	"app/model"
)

func GetUsers() []model.User{
	users := model.GetUsers()
	return users
}