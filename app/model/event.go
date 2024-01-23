package model

import (
	"time"
	"github.com/go-playground/validator/v10"
)

type Event struct {
	EventId uint `gorm:"primaryKey;autoIncrement"`
	Title string `validate:"required,lt=256"`
	Body string
	EventDate time.Time
	JoinDeadlineDate time.Time
	Capacity int `validate:"gte=0,lt=1000"`
}

func (event *Event) InsertEvent() error {
	validate := validator.New()
	err := validate.Struct(event)
	if err != nil {
		return err
	}
	result := Db.Create(event)
	if result.Error != nil {
		return result.Error
	}
	return nil
}