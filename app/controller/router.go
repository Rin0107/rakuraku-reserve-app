package controller

import(
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine{
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.GET("/api/admin/users", GetUsers)
	r.POST("/api/admin/users/create",CreateUsers)
	r.GET("/api/equipments/", GetEquipments)
	r.POST("api/login",Login)
	return r
}
