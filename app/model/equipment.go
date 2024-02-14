package model

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
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

type EquipmentReservation struct {
	EquipmentReservationId uint      `gorm:"primaryKey;autoIncrement"`
	UserId                 int       `validate:"required"`
	EquipmentId            int       `validate:"required"`
	ReservationStartTime   time.Time `validate:"required"`
	ReservationEndTime     time.Time `validate:"required"`
	ActivityStartTime      time.Time `validate:"required"`
	ActivityEndTime        time.Time `validate:"required"`
}

// HTTP JSONリクエストデータ定義
type EquipmentReservingRequest struct {
	UserId               int    `json:"userId"`
	ReservationStartTime string `json:"reservationStartTime"`
	ReservationEndTime   string `json:"reservationEndTime"`
	ActivityStartTime    string `json:"activityStartTime"`
	ActivityEndTime      string `json:"activityEndTime"`
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
