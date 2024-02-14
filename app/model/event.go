package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Event struct {
	EventId          uint   `gorm:"primaryKey;autoIncrement"`
	Title            string `validate:"required,lt=256"`
	Body             string
	EventDate        time.Time
	JoinDeadlineDate time.Time
	Capacity         int `validate:"gte=0,lt=1000"`
}

/*
データベースに新規イベントを挿入する。
バリデーションを実行し、有効なイベントか確認する。
バリデーションが失敗した場合、エラーを返す。
バリデーションが成功した場合、データベースにイベントを挿入を試みる。
イベントの挿入に失敗した場合、エラーを返す。
*/
func (event *Event) InsertEvent() error {
	// バリデーションを新規に作成
	validate := validator.New()

	// イベントをバリデーション
	err := validate.Struct(event)
	if err != nil {
		return err
	}

	// データベースにイベントを挿入
	result := Db.Create(event)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
