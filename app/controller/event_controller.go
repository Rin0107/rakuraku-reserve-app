package controller

import (
	"app/service"

	"github.com/gin-gonic/gin"
)

/*
新規イベントをデータベースに挿入するメソッド
リクエストコンテキストからイベント情報を取得し、サービス層を介してデータベースに挿入を試みる。
正常にイベントが挿入されたら、ステータスコード200と成功メッセージをJSONで返す。
イベントの挿入がエラーとなった場合は、ステータスコード404とエラーメッセージをJSONで返す。
*/
func InsertEvent(c *gin.Context) {
	err := service.InsertEvent(c)
	if err == nil {
		c.IndentedJSON(200, gin.H{"message": "Event inserted successfully"})
		return
	}
	c.IndentedJSON(404, gin.H{"error": err.Error()})
}
