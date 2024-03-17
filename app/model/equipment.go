package model

import (
	"errors"
	"fmt"
	"time"
	"gorm.io/gorm"
)

type Equipment struct {
	EquipmentId         int    `json:"equipmentId,omitempty"`
	Name                string `json:"name,omitempty"`
	Explanation         string `json:"explanation,omitempty"`
	EquipmentCategoryId int    `json:"equipmentCategoryId,omitempty"`
	EquipmentImg        string `json:"equipmentImg,omitempty"`
	IsAvailable         bool   `json:"IsAvaliable,omitempty"`
}

func GetEquipments() (equipments []Equipment) {
	//Table名を指定しない場合に、equipment単数型のテーブル名としてみなされているので。
	result := Db.Table("equipments").Where("is_available = true").Find(&equipments)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

// データベースに新規の機材予約を挿入する。
func (equipmentReservation *EquipmentReservation) InsertEquipmentReservation() error {
	result := Db.Create(equipmentReservation)
	if result.Error != nil {
		fmt.Print(result.Error)
		return fmt.Errorf("failed to insert equipment reservation.")
	}
	return nil
}

// 指定の機材予約を論理削除する。
func DeleteEquipmentReservation(equipmentId int, reserveId int) error {
	var equipmentReservation EquipmentReservation

	// 指定の機材予約の存在チェック
	if err := Db.Where("equipment_reservation_id = ? AND equipment_id = ? AND deleted_at IS NULL", reserveId, equipmentId).
		First(&equipmentReservation).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("equipment reservation not found")
		}
		fmt.Print(err)
		return fmt.Errorf("failed to find equipment reservation:")
	}

	if err := Db.Model(&equipmentReservation).Update("deleted_at", time.Now()).Error; err != nil {
		fmt.Print(err)
		return fmt.Errorf("failed to delete equipment reservation.")
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

// リクエストから受け取ったデータを使用し、機材予約を変更する
func (newEquipmentReservation *EquipmentReservation) ChangeEquipmentReservation() error {
	var equipmentReservationId = newEquipmentReservation.EquipmentReservationId
	var equipmentId = newEquipmentReservation.EquipmentId

	var existingReservation EquipmentReservation

	// 該当の機材予約の存在チェック
	if err := Db.Where("equipment_reservation_id = ? AND equipment_id = ? AND deleted_at IS NULL", equipmentReservationId, equipmentId).
		First(&existingReservation).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("equipment reservation not found")
		}
		fmt.Print(err)
		return fmt.Errorf("failed to find equipment reservation:")
	}

	// リクエストから受け取ったデータに更新
	existingReservation.ReservationStartTime = newEquipmentReservation.ActivityStartTime
	existingReservation.ReservationEndTime = newEquipmentReservation.ReservationEndTime
	existingReservation.ActivityStartTime = newEquipmentReservation.ActivityStartTime
	existingReservation.ActivityEndTime = newEquipmentReservation.ActivityEndTime

	// 更新をデータベースに反映
	if err := Db.Save(&existingReservation).
		Error; err != nil {
		fmt.Print(err)
		return fmt.Errorf("failed to delete equipment reservation.")
	}

	return nil
}
