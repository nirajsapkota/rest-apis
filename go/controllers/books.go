package controllers

import (
	"fmt"
	"net/http"

	"example.com/hello/models"
	"github.com/gin-gonic/gin"
)

var books = []models.Book{}
var count int = 0

func CreateBook(c *gin.Context) {
	var req models.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count += 1
	book := models.Book{ID: count, Title: req.Title, Author: req.Author}
	books = append(books, book)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Created Book with ID %v", count)})
}

func ReadBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func UpdateBook(c *gin.Context) {
	var req models.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range books {
		if books[i].ID == *(req.ID) {

			if req.Title != "" {
				books[i].Title = req.Title
			}

			if req.Author != "" {
				books[i].Author = req.Author
			}

			break

		}
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Updated Book with ID %v", *(req.ID))})
}

func DeleteBook(c *gin.Context) {
	var req models.DeleteBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books = append(books[:*(req.ID)-1], books[*(req.ID):]...)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Deleted Book with ID %v", *(req.ID))})
}
