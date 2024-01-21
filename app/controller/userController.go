package controller

import(
	"app/model"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context){
	users := model.GetUsers()
	switch len(users){
	case 0:
		c.IndentedJSON(404, users)
	default:
		c.IndentedJSON(200, users)
	}
}