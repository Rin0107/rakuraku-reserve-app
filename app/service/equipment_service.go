package service

import(
	"app/model"
)

func GetEquipments() []model.Equipment{
	equipments := model.GetEquipments()
	return equipments
}

// func GetEquipmentById(equipmentId int) model.Equipment{
// 	// equipment := model.GetDetailEquipment(equipmentId int)
// 	return 1;
// }
func GetEquipmentById(equipmentId int) model.Equipment{
	equipment := model.GetEquipmentById(equipmentId)
	return equipment;
}
