package model

import (
	"fmt"
	"time"
)

// "gorm.io/gorm"

type Equipment struct {
	// gorm.Model `json:"-"`
	EquipmentId         int
	Name                string
	Explanation         string
	EquipmentCategoryId int
	EquipmentImg        string
	IsAvailable         bool
}

// func GetAll() (equiments []Equiment) {
// 	result := Db.Find(&equiments)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }

// func (u *User) Create() {
// 	result := Db.Create(u)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }

type EquipmentReservation struct {
	EquipmentReservationId uint      `gorm:"primaryKey;autoIncrement"`
	UserId                 int       `validate:"required"`
	EquipmentId            int       `validate:"required"`
	ReservationStartTime   time.Time `validate:"required"`
	ReservationEndTime     time.Time `validate:"required"`
	ActivityStartTime      time.Time `validate:"required"`
	ActivityEndTime        time.Time `validate:"required"`
}

func GetEquipments() (equipments []Equipment) {
	//Table名を指定しない場合に、equipment単数型のテーブル名としてみなされているので。
	result := Db.Table("equipments").Where("is_available = true").Find(&equipments)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func DeleteEquipmentReservation(equipmentId int, reserveId int) error {
	result := Db.Delete(&EquipmentReservation{}, reserveId)
	if result.Error != nil {
		return result.Error
	}
	fmt.Print(result)
	return nil
}
