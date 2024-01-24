package controller

import (
	"app/service"
	"strconv"
	"github.com/gin-gonic/gin"
)

//機材一覧を取得するAPI（is_availableがtrue）
func GetEquipments(c *gin.Context){
	equipments := service.GetEquipments()
	switch len(equipments){
	case 0:
		c.IndentedJSON(404, equipments)
	default:
		c.IndentedJSON(200, equipments)
	}
}

//機材の詳細情報を取得(idにて取得)
func GetEquipmentById(c *gin.Context){
	strEquipmentId :=c.Param("equipmentId");
	//pathから受け取ったstring型をint型に変換
	equipmentId, _ := strconv.Atoi(strEquipmentId);
	equipment := service.GetEquipmentById(equipmentId);
	switch equipment.EquipmentId{
	case 0:
		c.IndentedJSON(404, equipment)
	default:
		c.IndentedJSON(200, equipment)
	}
}
