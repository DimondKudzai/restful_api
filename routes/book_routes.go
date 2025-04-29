package routes

// import impotant packages

import (
	"github.com/gin-gonic/gin"
	"../models"
	"log"
	"../utils"
)


// setup all routes functiom

var books []models.Book

func SetupRoutes(r *gin.RouterGroup) {
	bookGroup := r.Group("/books")
	{
		bookGroup.GET("/all_books", getBooks)
		bookGroup.GET("/:id", getBook)
		bookGroup.POST("/create", createBook)
		bookGroup.PUT("/:id", updateBook)
		bookGroup.DELETE("/:id", deleteBook)
	}
}


//get all books

func getBooks(c *gin.Context) {
	c.JSON(200, books)
}


// get book

func getBook(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(200, book)
			return
		}
	}
	c.JSON(404, gin.H{"message": "Book not found"})
}


// create book

func createBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(400, gin.H{"message": "Invalid
		 request"})
		return
	}
	books = append(books, newBook)
	c.JSON(201, newBook)
}



// update book

func updateBook(c *gin.Context) {
	id := c.Param("id")
	for index, book := range books {
		if book.ID == id {
			var updatedBook models.Book
			if err := c.BindJSON(&updatedBook); err != nil {
				c.JSON(400, gin.H{"message": "Invalid request"})
				return
			}
			books[index] = updatedBook
			c.JSON(200, updatedBook)
			return
		}
	}
	c.JSON(404, gin.H{"message": "Book not found"})
}



// delete book


func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			c.JSON(200, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"message": "Book not found"})
}