package controller

import (
	"app/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

//ユーザー登録フォーム情報
//name:空文字・0文字以上100文字以下
//email:メール形式・メール重複
//role:admin,userでの入力制限（仕様変更可能性あり）
type User struct {
	Name  string `json:"name" validate:"required,gte=0,lte=100"`
	Email string `json:"email" validate:"email,email_unique_validation"`
	Role  string `json:"role" validate:"oneof=admin user"`
}

//エラーメッセージ情報
type Error struct{
	Message string
}

var validate *validator.Validate

//管理者がユーザー登録するためのメソッド
//登録成功時、登録内容（name,email,role）を201で返す
//バリデーション時、エラーメッセージを400で返す
//その他DBエラー時等は500で返す
func CreateUsers(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 以下バリデーション実装
	validate = validator.New()
	validate.RegisterValidation("email_unique_validation",emailUniqueValidation)
	validateErr := validate.Struct(user)
	if validateErr!=nil{
		for _, err := range validateErr.(validator.ValidationErrors) {
			fmt.Println("Namespace =",err.Namespace())
			fmt.Println("Tag =",err.Tag())
			fmt.Println("Type =",err.Type())
			fmt.Println("Value =",err.Value())
			fmt.Println("Param =",err.Param())
		}
		var errorMessage Error
		errorMessage.Message="登録が失敗しました"
		c.IndentedJSON(400, errorMessage)
	}else{
		service.CreateUsers(user.Name,user.Email,user.Role)
		c.IndentedJSON(201, user)
	}
}

// ユーザーを理論削除するためのメソッド
func DeleteUser(c *gin.Context){
	// URLからuserIdを取得する
	userId,_ := strconv.Atoi(c.Param("userId"))
	err:=service.DeleteUser(userId)
	if err != nil {
		fmt.Println(err)
		errorMessage := ResponseMessage{Message: "ユーザー削除に失敗しました"}
		c.JSON(500,errorMessage)
		return
	}
	message := ResponseMessage{Message: "ユーザーが削除されました"}
	c.JSON(200,message)
}

//メール重複に関するカスタムバリデーション実装
func emailUniqueValidation(fl validator.FieldLevel) bool{
	email := fl.Field().String()
	return service.IsEmail(email)
}