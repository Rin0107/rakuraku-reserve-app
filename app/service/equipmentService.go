package service

import(
	"app/model"
)

func GetEquipments() []model.Equiments{
	equiments := model.GetEquipments()
	return equiments
}
