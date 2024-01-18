package controller

import(
	"github.com/gin-gonic/gin"
)

type User struct{
	Id int `json:"id`
	Name string `json:"name`
}

func Index(c *gin.Context){
	c.HTML(200, "index.html", gin.H{})
}

func GetUser(c *gin.Context){
	user := User{
		Id: 1,
		Name: "タロウ",
	}
	c.IndentedJSON(200, user)
}