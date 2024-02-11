package controller

import (
	"app/service"
	"strconv"
	"github.com/gin-gonic/gin"
)

//機材一覧を取得するAPI（is_availableがtrue）
func GetEquipmentReservationsByUserId(c *gin.Context){
	strUserId :=c.Param("userId");
	//pathから受け取ったstring型をint型に変換
	userId, e := strconv.Atoi(strUserId);
	equipmentReservations := service.GetEquipmentReservationsByUserId(userId)
	var message string
	if e != nil {
		message = "正しい数値を入力してください。"
	}
	switch len(equipmentReservations){
	case 0:
		c.IndentedJSON(404, gin.H{
			"status":404,
			"message": message,
			"equipmentReservations":equipmentReservations,
		})
	default:
		c.IndentedJSON(200, gin.H{
			"status":202,
			"message": message,
			"equipmentReservations":equipmentReservations,
		})
	}
}

//機材の詳細情報を取得(idにて取得)
// func GetEquipmentById(c *gin.Context){
// 	strEquipmentId :=c.Param("equipmentId");
// 	//pathから受け取ったstring型をint型に変換
// 	equipmentId, e := strconv.Atoi(strEquipmentId);
	
// 	equipment := service.GetEquipmentById(equipmentId);
	
// 	var message string
// 	if e != nil {
// 		message = "機材IDには正しい数値を入力してください"
// 	}else if equipment.EquipmentId == 0 {
// 		message = "機材が存在しません"
// 	}
// 	switch equipment.EquipmentId{
// 	case 0:
// 		c.IndentedJSON(404, gin.H{
// 			"status":"404",
// 			"equipment": equipment,
// 			"message": message,
// 		})
// 	default:
// 		c.IndentedJSON(200, gin.H{
// 			"status":"200",
// 			"equipment": equipment,
// 			"message": message,
// 		})
// 	}
// }
