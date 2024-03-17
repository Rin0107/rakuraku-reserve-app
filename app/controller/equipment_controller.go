package controller

import (
	"app/request"
	"app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EquipmentServiceImplementation struct{}

/*
機材予約を論理削除するメソッド
正常に更新されたら、ステータスコード200と成功メッセージをJSONで返す。
更新がエラーとなった場合は、ステータスコード404とエラーメッセージをJSONで返す。
*/
func DeleteEquipmentReservation(c *gin.Context) {
	equipmentId := c.Param("equipmentId")
	reserveId := c.Param("reserveId")

	err := service.DeleteEquipmentReservation(equipmentId, reserveId)
	if err != nil {
		c.IndentedJSON(404, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"message": "Reserved equipment successfully"})
}

// 機材一覧を取得するAPI（is_availableがtrue）
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
	equipmentId := c.Param("equipmentId")
	var equipmentReservingRequest request.EquipmentReservingRequest

	if err := c.ShouldBindJSON(&equipmentReservingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
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

// 機材の詳細情報を取得(idにて取得)
func GetEquipmentById(c *gin.Context) {
	strEquipmentId := c.Param("equipmentId")
	//pathから受け取ったstring型をint型に変換
	equipmentId, e := strconv.Atoi(strEquipmentId)

	equipment := service.GetEquipmentById(equipmentId)

	var message string
	if e != nil {
		message = "機材IDには正しい数値を入力してください"
	} else if equipment.EquipmentId == 0 {
		message = "機材が存在しません"
	}
	switch equipment.EquipmentId {
	case 0:
		c.IndentedJSON(404, gin.H{
			"status":    "404",
			"equipment": equipment,
			"message":   message,
		})
	default:
		c.IndentedJSON(200, gin.H{
			"status":    "200",
			"equipment": equipment,
			"message":   message,
		})
	}
}

/*
機材予約を変更するメソッド
リクエストコンテキストから機材予約情報を取得し、サービス層を介してデータベースの更新を試みる。
正常に更新されたら、ステータスコード200と成功メッセージをJSONで返す。
更新がエラーとなった場合は、ステータスコード404とエラーメッセージをJSONで返す。
*/
func changeEquipmentReservation(c *gin.Context) {
	equipmentId := c.Param("equipmentId")
	reserveId := c.Param("reserveId")

	var equipmentReservingRequest request.EquipmentReservingRequest

	if err := c.ShouldBindJSON(&equipmentReservingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(equipmentReservingRequest)
	if err != nil {
		c.IndentedJSON(404, gin.H{"error": err.Error()})
		return
	}

	err = service.ChangeEquipmentReservation(equipmentId, reserveId, equipmentReservingRequest)
	if err != nil {
		c.IndentedJSON(404, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"message": "Changed equipment reservation successfully"})
}
