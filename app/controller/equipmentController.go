package controller

import (
	"app/service"

	"github.com/gin-gonic/gin"
)

func GetEquipments(c *gin.Context){
	equiments := service.GetEquipments()
	switch len(equiments){
	case 0:
		c.IndentedJSON(404, equiments)
	default:
		c.IndentedJSON(200, equiments)
	}
}
