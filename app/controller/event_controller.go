package controller

import (
	"app/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTP JSONリクエストデータ定義
type NewEventInfo struct {
	Title            string `json:"title"`
	Body             string `json:"body"`
	EventDate        string `json:"eventDate"`
	JoinDeadlineDate string `json:"joinDeadlineDate"`
	Capacity         string `json:"capacity"`
}

/*
新規イベントをデータベースに挿入するメソッド
リクエストコンテキストからイベント情報を取得し、サービス層を介してデータベースに挿入を試みる。
正常にイベントが挿入されたら、ステータスコード200と成功メッセージをJSONで返す。
イベントの挿入がエラーとなった場合は、ステータスコード404とエラーメッセージをJSONで返す。
*/
func InsertEvent(c *gin.Context) {
	var newEventInfo NewEventInfo
	if err := c.ShouldBindJSON(&newEventInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print(newEventInfo)
	err := service.InsertEvent(newEventInfo.Title, newEventInfo.Body, newEventInfo.EventDate, newEventInfo.JoinDeadlineDate, newEventInfo.Capacity)
	if err == nil {
		c.IndentedJSON(200, gin.H{"message": "Event inserted successfully"})
		return
	}
	c.IndentedJSON(404, gin.H{"error": err.Error()})
}
