package model

import (
	"time"
	"gorm.io/gorm"
	// "github.com/go-playground/validator/v10"
)

//機材予約の構造体
type EquipmentReservation struct {
	EquipmentReservationId uint `gorm:"primaryKey;autoIncrement" ,json:"equipmentReservationId,omitempty"`
	UserId int `json:"userId,omitempty"`
	EquipmentId int `json:"equipmentId,omitempty"`
	ReservationStartTime time.Time `json:"reservationStartTime,omitempty"`
	ReservationEndTime time.Time `json:"reservationEndTime,omitempty"`
	ActivityStartTime time.Time `json:"activityStartTime,omitempty"`
	ActivityEndTime time.Time `json:"activityEndTime,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

//予約をしているユーザーの一部情報
type WithUser struct {
	UserId int `gorm:"primaryKey;autoIncrement"`
	Name string
	UserIcon string
}

//機材情報の一部
type WithEquipment struct {
	EquipmentId int `gorm:"primaryKey;autoIncrement"`
	Name string 
}

//機材予約の構造体
type EquipmentReservationWithUserAndEquipment struct {
	EquipmentReservationId uint `gorm:"primaryKey;autoIncrement" ,json:"equipmentReservationId,omitempty"`
	UserId int `json:"userId,omitempty"`
	EquipmentId int `json:"equipmentId,omitempty"`
	ReservationStartTime time.Time `json:"reservationStartTime,omitempty"`
	ReservationEndTime time.Time `json:"reservationEndTime,omitempty"`
	ActivityStartTime time.Time `json:"activityStartTime,omitempty"`
	ActivityEndTime time.Time `json:"activityEndTime,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Equipment WithEquipment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;",json:"Equipment,omitempty"`
	User WithUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;",json:"User,omitempty"`
}

var equipment Equipment
var equipmentReservations []EquipmentReservation 
var equipmentReservationsWithUserAndEquipment []EquipmentReservationWithUserAndEquipment

//特定の機材に関する予約情報一覧を取得
func GetEquipmentReservationsByEquipmentId(equipmentId int)(equipmentReservationsWithUserAndEquipment []EquipmentReservationWithUserAndEquipment){
	result := Db.Debug().Model(&EquipmentReservation{}).Preload("User", func(Db *gorm.DB) *gorm.DB {
			return Db.Model(&User{})// 見に行くテーブルはusersであると明示
		}).Preload("Equipment", func(Db *gorm.DB)*gorm.DB{
			return Db.Table("equipments").Find(&Equipment{})
	}).Where("equipment_id = ?",equipmentId).Where("deleted_at is null").Where("reservation_end_time > ?",time.Now()).Order("reservation_start_time asc").Find(&equipmentReservationsWithUserAndEquipment)
	
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

// 特定のユーザーが予約をしている情報一覧を取得
func GetEquipmentReservationsByUserId(userId int)(equipmentReservationsWithUserAndEquipment []EquipmentReservationWithUserAndEquipment){
	result := Db.Debug().Model(&EquipmentReservation{}).Preload("User", func(Db *gorm.DB) *gorm.DB {
				return Db.Model(&User{})// 見に行くテーブルはusersであると明示
			}).Preload("Equipment", func(Db *gorm.DB) *gorm.DB {
				return Db.Table("equipments").Find(&Equipment{})// 見に行くテーブルはequipmentであると明示
			}).Where("user_id = ?",userId).Where("deleted_at is null").Where("reservation_end_time > ?",time.Now()).Order("reservation_start_time asc").Find(&equipmentReservationsWithUserAndEquipment)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
