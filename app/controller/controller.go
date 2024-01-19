package controller

import(
	"app/model"

	"github.com/gin-gonic/gin"
)

type User struct{
	Id int `json:"id`
	Name string `json:"name`
}

func Index(c *gin.Context){
	users := model.GetAll()
	c.HTML(200, "index.html", gin.H{"users": users})
}

func GetUser(c *gin.Context){
	user := User{
		Id: 1,
		Name: "タロウ",
	}
	c.IndentedJSON(200, user)
}

func PostCreate(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := model.User{Name: name, Email: email, Password: password}
	user.Create()
	c.IndentedJSON(200, user)
	// c.Redirect(301, "/")
}