package service

import(
	"app/model"
)

func GetEquipments() []model.Equipment{
	equipments := model.GetEquipments()
	return equipments
}
