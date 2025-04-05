package delivery

import (
	"net/http"
	"strconv"

	"bookProject/internal/models"
	"bookProject/internal/services"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(bookService services.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := bookService.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch books"})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}

func GetBookByID(bookService services.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		book, err := bookService.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

func CreateBook(bookService services.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		if err := bookService.Create(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create book"})
			return
		}
		c.JSON(http.StatusCreated, book)
	}
}

func UpdateBook(bookService services.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		book.ID = uint(id)
		if err := bookService.Update(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update book"})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

func DeleteBook(bookService services.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := bookService.Delete(uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete book"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
