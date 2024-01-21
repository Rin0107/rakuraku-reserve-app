package controller

import (
	"app/service"

	"github.com/gin-gonic/gin"
)

func GetEquipments(c *gin.Context){
	equipments := service.GetEquipments()
	switch len(equipments){
	case 0:
		c.IndentedJSON(404, equipments)
	default:
		c.IndentedJSON(200, equipments)
	}
}
