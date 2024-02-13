package controller

import (
	"app/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type EquipmentReservationInfo struct {
	UserId               int    `json:"userId"`
	ReservationStartTime string `json:"reservationStartTime"`
	ReservationEndTime   string `json:"reservationEndTime"`
	ActivityStartTime    string `json:"activityStartTime"`
	ActivityEndTime      string `json:"activityEndTime"`
}

func GetEquipments(c *gin.Context) {
	equipments := service.GetEquipments()
	switch len(equipments) {
	case 0:
		c.IndentedJSON(404, equipments)
	default:
		c.IndentedJSON(200, equipments)
	}
}

func ReserveEquipment(c *gin.Context) {
	id := c.Param("equipmentId")
	fmt.Print(id)
}
