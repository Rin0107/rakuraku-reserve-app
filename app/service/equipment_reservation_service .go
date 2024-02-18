package service

import(
	"app/model"
)
/*特定の機材の機材予約一覧を取得する*/
func GetEquipmentReservationsByEquipmentId(equipmentId int) []model.EquipmentReservation{
	equipmentReservations := model.GetEquipmentReservationsByEquipmentId(equipmentId)
	return equipmentReservations
}
/*特定のユーザーの機材予約一覧を取得する*/
func GetEquipmentReservationsByUserId(userId int) []model.EquipmentReservation{
	equipmentReservations := model.GetEquipmentReservationsByUserId(userId)
	return equipmentReservations
}
