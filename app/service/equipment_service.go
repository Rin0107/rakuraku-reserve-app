package service

import (
	"app/model"
	"app/request"
	"strconv"
	"time"
)

func GetEquipments() []model.Equipment {
	equipments := model.GetEquipments()
	return equipments
}

/*
リクエストから受け取ったデータの型を変換し、新規機材予約をデータベースに挿入する。
正常に変換処理ができなかった場合は、エラーを返す。
equipmentReservationモデルにリクエストから受け取ったデータを渡し、
モデルのInsertEquipmentReservationメソッドを呼び出す。
*/
func ReserveEquipment(equipmentIdStr string, reservingRequest request.EquipmentReservingRequest) error {
	// equipmentId(string)をintに変換
	equipmentId, err := strconv.Atoi(equipmentIdStr)
	if err != nil {
		return err
	}

	// 各時間の変換
	reservationStartTime, err := StrToTime(reservingRequest.ReservationStartTime)
	if err != nil {
		return err
	}
	reservationEndTime, err := StrToTime(reservingRequest.ReservationEndTime)
	if err != nil {
		return err
	}
	activityStartTime, err := StrToTime(reservingRequest.ActivityStartTime)
	if err != nil {
		return err
	}
	activityEndTime, err := StrToTime(reservingRequest.ActivityEndTime)
	if err != nil {
		return err
	}

	equipmentReservation := model.EquipmentReservation{
		UserId:               reservingRequest.UserId,
		EquipmentId:          equipmentId,
		ReservationStartTime: reservationStartTime,
		ReservationEndTime:   reservationEndTime,
		ActivityStartTime:    activityStartTime,
		ActivityEndTime:      activityEndTime,
	}

	err = equipmentReservation.InsertEquipmentReservation()
	if err != nil {
		return err
	}

	// 処理の中でエラーが発生しなかった場合、nilを返す。
	return nil
}

/*
パラメータで受け取ったIDの機材予約を論理削除するため、
equipment_reservationモデルのDeleteEquipmentReservationメソッドを呼び出す。
*/
func DeleteEquipmentReservation(equipmentIdStr string, reserveIdStr string) error {
	// equipmentId(string)をintに変換
	equipmentId, err := strconv.Atoi(equipmentIdStr)
	if err != nil {
		return err
	}
	// reserveId(string)をintに変換
	reserveId, err := strconv.Atoi(reserveIdStr)
	if err != nil {
		return err
	}

	err = model.DeleteEquipmentReservation(equipmentId, reserveId)
	if err != nil {
		return err
	}

	// 処理の中でエラーが発生しなかった場合、nilを返す。
	return nil
}

// 日時(string)をtime.Time型に変換する関数
func StrToTime(dateStr string) (time.Time, error) {
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

/*機材をIDから1つだけ取得する。*/
func GetEquipmentById(equipmentId int) model.Equipment {
	equipment := model.GetEquipmentById(equipmentId)
	return equipment
}
