package controller

import (
	"app/service"
	"strconv"
	"github.com/gin-gonic/gin"
)

//機材IDに紐づく予約情報を取得する。
func GetEquipmentReservationsByEquipmentId(c *gin.Context){
	strEquipmentId :=c.Param("equipmentId");
	//pathから受け取ったstring型をint型に変換
	equipmentId, e := strconv.Atoi(strEquipmentId);
	equipmentReservations := service.GetEquipmentReservationsByEquipmentId(equipmentId)
	var message string
	if e != nil {
		c.IndentedJSON(404, gin.H{
			"status":404,
			"message": "正しい数値を入力してください",
			"equipmentReservations":equipmentReservations,
		})
		return
	}
	switch len(equipmentReservations){
	case 0:
		c.IndentedJSON(404, gin.H{
			"status":404,
			"message": "予約が一件もございません",
			"equipmentReservations":equipmentReservations,
		})
	default:
		c.IndentedJSON(200, gin.H{
			"status":200,
			"message": message,
			"equipmentReservations":equipmentReservations,
		})
	}
}
// ユーザーIDに紐づく予約情報を取得する。
// func GetEquipmentReservationsByUserId(c *gin.Context){
// 	strUserId :=c.Param("userId");
// 	//pathから受け取ったstring型をint型に変換
// 	userId, e := strconv.Atoi(strUserId);
// 	equipmentReservations := service.GetEquipmentReservationsByUserId(userId)
// 	var message string
// 	if e != nil {
// 		c.IndentedJSON(404, gin.H{
// 			"status":404,
// 			"message": "正しい数値を入力してください",
// 			"equipmentReservations":equipmentReservations,
// 		})
// 		return
// 	}
// 	switch len(equipmentReservations){
// 	case 0:
// 		c.IndentedJSON(404, gin.H{
// 			"status":404,
// 			"message": "予約が一件もございません",
// 			"equipmentReservations":equipmentReservations,
// 		})
// 	default:
// 		c.IndentedJSON(200, gin.H{
// 			"status":200,
// 			"message": message,
// 			"equipmentReservations":equipmentReservations,
// 		})
// 	}
// }
