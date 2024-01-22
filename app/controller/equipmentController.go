package controller

import (
	"app/service"
	"strconv"
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

func GetEquipmentById(c *gin.Context){
	strEquipmentId :=c.Param("equipmentId")
	equipmentId, _ := strconv.Atoi(strEquipmentId)
	equipment := service.GetEquipmentById(equipmentId)
	c.IndentedJSON(200, equipment)
	// switch (equipment){
	// case false:
	// 	c.IndentedJSON(404, equipment)
	// default:
	// 	c.IndentedJSON(200, equipment)
	// }
}
