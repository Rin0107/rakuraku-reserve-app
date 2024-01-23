package controller

import (
	"app/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type User struct {
	Name  string `json:"name" validate:"required,gte=0,lte=100"`
	Email string `json:"email" validate:"required,email"`
	Role  string `json:"role" validate:"oneof=admin user"`
}

var validate *validator.Validate

func CreateUsers(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate = validator.New()
	validateErr := validate.Struct(user)
	if validateErr!=nil{
		for _, err := range validateErr.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		c.IndentedJSON(400, user)
	}
	service.CreateUsers(user.Name,user.Email,user.Role)
	c.IndentedJSON(201, user)
}