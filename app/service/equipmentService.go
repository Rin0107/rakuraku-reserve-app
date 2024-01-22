package service

import(
	"app/model"
)

func GetEquipments() []model.Equipments{
	equipments := model.GetEquipments()
	return equipments
}
