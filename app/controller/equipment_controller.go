package controller

import (
	"app/service"
	"strconv"
	"github.com/gin-gonic/gin"
	// "fmt"
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
	equipmentId, e := strconv.Atoi(strEquipmentId);
	var message string
	
	
	if e != nil {
		message = "機材IDには正しい数値を入力してください"
	}
	equipment := service.GetEquipmentById(equipmentId);
	switch equipment.EquipmentId{
	case 0:
		c.IndentedJSON(404, equipment)
	default:
		c.IndentedJSON(200, equipment)
	}
}
