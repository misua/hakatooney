package main

import (
	"github.com/gin-gonic/gin"
	"hakatooney/controllers"
	"hakatooney/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)

	r.Run()

}