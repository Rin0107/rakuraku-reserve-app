package service

import(
	"app/model"
)
/*機材一覧を取得する（is_availbleがtrue）*/
func GetEquipments() []model.Equipment{
	equipments := model.GetEquipments()
	return equipments
}

// func GetEquipmentById(equipmentId int) model.Equipment{
// 	// equipment := model.GetDetailEquipment(equipmentId int)
// 	return 1;
// }

/*機材をIDから1つだけ取得する。*/
func GetEquipmentById(equipmentId int) model.Equipment{
	equipment := model.GetEquipmentById(equipmentId)
	return equipment;
}
