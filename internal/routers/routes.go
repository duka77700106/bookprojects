package routers

import (
	"bookProject/internal/delivery"
	"bookProject/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, bookService services.BookService) {
	bookRoutes := router.Group("/books")
	{
		bookRoutes.GET("/", delivery.GetAllBooks(bookService))
		bookRoutes.GET("/:id", delivery.GetBookByID(bookService))
		bookRoutes.POST("/", delivery.CreateBook(bookService))
		bookRoutes.PUT("/:id", delivery.UpdateBook(bookService))
		bookRoutes.DELETE("/:id", delivery.DeleteBook(bookService))
	}
}
