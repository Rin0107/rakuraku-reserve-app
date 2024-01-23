package main

import "app/controller"

// "github.com/gin-gonic/gin"

func main(){
	router := controller.GetRouter()
	router.Run(":8080")
}