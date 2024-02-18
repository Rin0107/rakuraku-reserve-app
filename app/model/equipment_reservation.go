package model

import (
	"time"
	// "gorm.io/gorm"
	// "github.com/go-playground/validator/v10"
)

type EquipmentReservation struct {
	EquipmentReservationId int `json:"equipmentReservationId,omitempty"`
	UserId int `json:"userId,omitempty"`
	EquipmentId int `json:"equipmentId,omitempty"`
	ReservationStartTime time.Time `json:"reservationStartTime,omitempty"`
	ReservationEndTime time.Time `json:"reservationEndTime,omitempty"`
	ActivityStartTime time.Time `json:"activityStartTime,omitempty"`
	ActivityEndTime time.Time `json:"activityEndTime,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}


//特定の機材に関する予約情報一覧を取得
func GetEquipmentReservationsByEquipmentId(equipmentId int)(equipmentReservations []EquipmentReservation){
	//Table名を指定しない場合に、equipment単数型のテーブル名としてみなされているので。
	
	result := Db.Where("equipment_id = ?",equipmentId).Where("deleted_at is null").Where("reservation_end_time > ?",time.Now()).Order("equipment_id asc").Find(&equipmentReservations);
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

//特定のユーザーが予約をしている情報一覧を取得
func GetEquipmentReservationsByUserId(userId int)(equipmentReservations []EquipmentReservation){
	//Table名を指定しない場合に、equipment単数型のテーブル名としてみなされているので。
	
	result := Db.Table("equipment_reservations").Where("user_id = ?",userId).Where("deleted_at is null").Where("reservation_end_time > ?",time.Now()).Find(&equipmentReservations)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
