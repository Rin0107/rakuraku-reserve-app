package service

import (
	"app/model"
	"strconv"
	"time"
)

/*
リクエストから受け取ったデータを使用し、新規イベントをデータベースに挿入する。
リクエストからタイトル、本文、イベント日時、参加締切日時、定員を取得し、
正常に取得できない場合はエラーを返す。
取得したデータを使用して新しいEventモデルを作成し、モデルのInsertEventメソッドを呼び出す。
*/
func InsertEvent(title string, body string, eventDateStr string, joinDeadlineDateStr string, capacityStr string) error {

	// イベント日時をstringからtime.Time型に変換
	eventDate, err := time.Parse(time.RFC3339, eventDateStr)
	if err != nil {
		return err
	}

	// 参加締切日時をstringからtime.Time型に変換
	joinDeadlineDate, err := time.Parse(time.RFC3339, joinDeadlineDateStr)
	if err != nil {
		return err
	}

	// 定員をstringから整数に変換
	capacity, err := strconv.Atoi(capacityStr)
	if err != nil {
		return err
	}

	event := model.Event{
		Title:            title,
		Body:             body,
		EventDate:        eventDate,
		JoinDeadlineDate: joinDeadlineDate,
		Capacity:         capacity,
	}

	err = event.InsertEvent()
	if err != nil {
		return err
	}

	// エラーが発生しなかった場合、nilを返す。
	return nil
}
