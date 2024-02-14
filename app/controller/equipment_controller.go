package controller

import (
	"app/service"

	"github.com/gin-gonic/gin"
)

type EquipmentServiceImplementation struct{}

// func GetEquipments(c *gin.Context) {
// 	equipments := service.GetEquipments()
// 	switch len(equipments) {
// 	case 0:
// 		c.IndentedJSON(404, equipments)
// 	default:
// 		c.IndentedJSON(200, equipments)
// 	}
// }

/*
機材予約を物理削除するメソッド
正常に挿入されたら、ステータスコード200と成功メッセージをJSONで返す。
挿入がエラーとなった場合は、ステータスコード404とエラーメッセージをJSONで返す。
*/
func DeleteEquipmentReservation(c *gin.Context) {
	equipmentId := c.Param("equipmentId")
	reserveId := c.Param("reserveId")

	err := service.DeleteEquipmentReservation(equipmentId, reserveId)
	if err == nil {
		c.IndentedJSON(200, gin.H{"message": "Reserved equipment successfully"})
		return
	}
	c.IndentedJSON(404, gin.H{"error": err.Error()})
}
