package controller

import(
	"app/model"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context){
	users := model.GetUsers()
	c.IndentedJSON(200, users)
}