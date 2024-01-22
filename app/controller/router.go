package controller

import(
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine{
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.GET("/", Index)
	r.GET("/user", GetUser)
	r.GET("/api/equipments/", GetEquipments)
	r.GET("/api/equipments/:equipmentId", GetEquipmentById)
	r.POST("/create/user", PostCreate)
	r.GET("/api/admin/users", GetUsers)
	return r
}
