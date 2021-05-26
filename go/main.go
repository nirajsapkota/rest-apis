package main

import (
	"example.com/hello/controllers"
	"github.com/gin-gonic/gin"
)

func setupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/books", controllers.ReadBooks)
	r.POST("/book", controllers.CreateBook)
	r.PUT("/book", controllers.UpdateBook)
	r.DELETE("/book", controllers.DeleteBook)

	return r
}

func main() {
	setupServer().Run()
}
