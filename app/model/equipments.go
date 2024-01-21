package model

import (
	"time"

	"gorm.io/gorm"
)

type Equpment struct {
	gorm.Model `json:"-"`
	EqupmentId int
	Name string
	Explanation string
	EquipmentCategoryId int
	EqupmentImg string
	IsAvailable bool
}

func GetAll() (equiments []Equiment) {
	result := Db.Find(&equiments)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

// func (u *User) Create() {
// 	result := Db.Create(u)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }

func GetEquipments()(equiments []Equiment){
	result := Db.Where("is_available IS true").Find(&equiments)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
