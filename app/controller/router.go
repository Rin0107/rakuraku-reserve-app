package controller

import(
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine{
	r := gin.Default()
	r.LoadHTMLGlob("views/*html")
	r.GET("/", Index)
	r.GET("/user", GetUser)
	return r
}