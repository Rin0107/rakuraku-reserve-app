package controller

import (
	"net/http"
	"app/service"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func CreateUsers(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.CreateUsers(user.Name,user.Email,user.Role)
	c.IndentedJSON(201, user)
}