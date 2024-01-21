package model

import (
	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model `json:"-"`
	EquipmentId int
	Name string
	Explanation string
	EquipmentCategoryId int
	EquipmentImg string
	IsAvailable bool
}

// func GetAll() (equiments []Equiment) {
// 	result := Db.Find(&equiments)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }

// func (u *User) Create() {
// 	result := Db.Create(u)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }

func GetEquipments()(equipments []Equipment){
	result := Db.Where("is_available IS true").Find(&equipments)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
