package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.GET("/api/equipments/", GetEquipments)
	r.POST("/api/events", InsertEvent)
	r.GET("/api/admin/users", CheckAuth, GetUsers)
	r.POST("/api/admin/users/create", CreateUsers)
	r.POST("/api/login", Login)
	r.POST("/api/logout", Logout)
	return r
}

// 各処理の前に認可情報を確認するためのメソッド
// GetRouterの各メソッドのパス後の引数として本メソッドを追加することで認可処理の対象となる
// /api/adminが含まれている場合はadmin権限が必要
// user権限の場合はセッションIDを保持していることのみを認可条件としている
func CheckAuth(c *gin.Context) {
	// セッションIDがなかった場合403を返し、処理が中断される
	sessionId, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(403, "session情報がありません")
		c.Abort()
		return
	}
	// サーバーに保存されているセッション情報とリクエストされたセッション情報を比較する
	// admin権限を持っている場合/api/adminを含むパスを認可する
	// 権限がなかった場合403を返し、処理が中断される
	sessionMutex.Lock()
	userInformation := sessions[sessionId]
	if strings.HasPrefix(c.Request.URL.Path, "/api/admin") && userInformation.Role != "admin" {
		c.JSON(403, "権限が不足しています")
		c.Abort()
		sessionMutex.Unlock()
		return
	}
	sessionMutex.Unlock()

	// 認可が成功されたため、後続の処理を実行する
	c.Next()
}
