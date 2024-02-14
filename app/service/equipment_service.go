package service

import (
	"app/model"
	"strconv"
)

// func GetEquipments() []model.Equipment {
// 	equipments := model.GetEquipments()
// 	return equipments
// }

/*
リクエストから受け取ったデータの型を変換し、新規機材予約をデータベースに挿入する。
正常に変換処理ができなかった場合は、エラーを返す。
equipmentReservationモデルにリクエストから受け取ったデータを渡し、
モデルのInsertEquipmentReservationメソッドを呼び出す。
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
