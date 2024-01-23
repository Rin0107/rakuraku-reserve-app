package controller

import (
	"app/service"

	"github.com/gin-gonic/gin"
)

func InsertEvent(c *gin.Context) {
	err := service.InsertEvent(c)
	if err == nil {
		c.IndentedJSON(200, gin.H{"message": "Event inserted successfully"})
		return
	}
	c.IndentedJSON(404, gin.H{"error": err.Error()})
}