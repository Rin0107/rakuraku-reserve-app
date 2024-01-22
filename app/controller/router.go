package controller

import(
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine{
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.GET("/api/admin/users", GetUsers)
	return r
}