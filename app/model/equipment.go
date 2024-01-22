package model

import (
	// "gorm.io/gorm"
)

type Equipment struct {
	// gorm.Model `json:"-"`
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
	//Table名を指定しない場合に、equipment単数型のテーブル名としてみなされているので。
	result := Db.Table("equipments").Where("is_available = true").Find(&equipments)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
func GetEquipmentById(equipmentId int)(equipment Equipment){
	//Table名を指定しない場合に、equipment単数型のテーブル名としてみなされているので。
	result := Db.Table("equipments").Where("is_available = true AND equipment_id = ?",equipmentId).First(&equipment)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
