package controller

import (
	"app/service"

	"github.com/gin-gonic/gin"
)

// ユーザー一覧を取得するためのメソッド
// ユーザーが存在する場合はユーザー情報をリストにして200で返す
// ユーザーが存在しない場合は空のリストを404で返す
func GetUsers(c *gin.Context){
	users := service.GetUsers()
	switch len(users){
	case 0:
		c.IndentedJSON(404, users)
	default:
		c.IndentedJSON(200, users)
	}
}