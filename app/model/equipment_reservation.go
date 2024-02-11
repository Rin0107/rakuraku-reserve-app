package model

import (
	"time"
	// "gorm.io/gorm"
	// "github.com/go-playground/validator/v10"
)

type EquipmentReservation struct {
	EquipmentReservationId int `json:"equipmentReservationId,omitempty"`
	UserId int `json:"userId,omitempty"`
	ReservationStartTime time.Time `json:"reservationStartTime,omitempty"`
	ReservationEndTime time.Time `json:"reservationEndTime,omitempty"`
	ActivityStartTime time.Time `json:"activityStartTime,omitempty"`
	ActivityEndTime time.Time `json:"activityEndTime,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

func GetEquipmentReservationsByUserId(userId int)(equipmentReservations []EquipmentReservation){
	//Table名を指定しない場合に、equipment単数型のテーブル名としてみなされているので。
	result := Db.Table("equipment_reservations").Where("user_id = ?",userId).Find(&equipmentReservations)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

// func GetEquipmentById(equipmentId int)(equipment Equipment){
// 	//Table名を指定しない場合に、equipment単数型のテーブル名としてみなされているので。
// 	result := Db.Table("equipments").Where("is_available = true AND equipment_id = ?",equipmentId).Find(&equipment)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }
