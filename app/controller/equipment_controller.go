package controller

import (
	"app/request"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EquipmentServiceImplementation struct{}

func GetEquipments(c *gin.Context) {
	equipments := service.GetEquipments()
	switch len(equipments) {
	case 0:
		c.IndentedJSON(404, equipments)
	default:
		c.IndentedJSON(200, equipments)
	}
}

/*
新規機材予約をデータベースに挿入するメソッド
リクエストコンテキストからイベント情報を取得し、サービス層を介してデータベースに挿入を試みる。
正常に挿入されたら、ステータスコード200と成功メッセージをJSONで返す。
挿入がエラーとなった場合は、ステータスコード404とエラーメッセージをJSONで返す。
*/
func ReserveEquipment(c *gin.Context) {
	validate := validator.New()
	equipmentId := c.Param("equipmentId")
	var equipmentReservingRequest request.EquipmentReservingRequest

	if err := c.ShouldBindJSON(&equipmentReservingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validate.Struct(equipmentReservingRequest)
	if err != nil {
		c.IndentedJSON(404, gin.H{"error": err.Error()})
		return
	}

	err = service.ReserveEquipment(equipmentId, equipmentReservingRequest)
	if err != nil {
		c.IndentedJSON(404, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"message": "Reserved equipment successfully"})
}
