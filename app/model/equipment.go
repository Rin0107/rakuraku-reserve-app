package model

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

// "gorm.io/gorm"

type Equipment struct {
	// gorm.Model `json:"-"`
	EquipmentId         int    `json:"equipmentId,omitempty"`
	Name                string `json:"name,omitempty"`
	Explanation         string `json:"explanation,omitempty"`
	EquipmentCategoryId int    `json:"equipmentCategoryId,omitempty"`
	EquipmentImg        string `json:"equipmentImg,omitempty"`
	IsAvailable         bool   `json:"IsAvaliable,omitempty"`
}

type EquipmentReservation struct {
	EquipmentReservationId uint `gorm:"primaryKey;autoIncrement"`
	UserId                 int
	EquipmentId            int
	ReservationStartTime   time.Time
	ReservationEndTime     time.Time
	ActivityStartTime      time.Time
	ActivityEndTime        time.Time
}

func GetEquipments() (equipments []Equipment) {
	//Table名を指定しない場合に、equipment単数型のテーブル名としてみなされているので。
	result := Db.Table("equipments").Where("is_available = true").Find(&equipments)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

/*
バリデーションを実行し、データベースに新規の機材予約を挿入する。
処理に失敗した場合、エラーを返す。
*/
func (equipmentReservation *EquipmentReservation) InsertEquipmentReservation() error {
	validate := validator.New()

	err := validate.Struct(equipmentReservation)
	if err != nil {
		return err
	}
	result := Db.Create(equipmentReservation)
	if result.Error != nil {
		fmt.Print(result.Error)
		return fmt.Errorf("failed to insert equipment reservation.")
	}
	return nil
}

func GetEquipmentById(equipmentId int) (equipment Equipment) {
	//Table名を指定しない場合に、equipment単数型のテーブル名としてみなされているので。
	result := Db.Table("equipments").Where("is_available = true AND equipment_id = ?", equipmentId).Find(&equipment)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
