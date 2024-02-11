package service

import(
	"app/model"
)
/*機材一覧を取得する（is_availbleがtrue）*/
func GetEquipmentReservationsByUserId(userId int) []model.EquipmentReservation{
	equipmentReservations := model.GetEquipmentReservationsByUserId(userId)
	return equipmentReservations
}

// /*機材をIDから1つだけ取得する。*/
// func GetEquipmentById(equipmentId int) model.Equipment{
// 	equipment := model.GetEquipmentById(equipmentId)
// 	return equipment;
// }
