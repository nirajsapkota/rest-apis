package main

import (
	"example.com/hello/controllers"
	"github.com/gin-gonic/gin"
)

func setupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/api/v1/books", controllers.ReadBooks)
	r.POST("/api/v1/book", controllers.CreateBook)
	r.PUT("/api/v1/book", controllers.UpdateBook)
	r.DELETE("/api/v1/book", controllers.DeleteBook)

	return r
}

func main() {
	setupServer().Run()
}
