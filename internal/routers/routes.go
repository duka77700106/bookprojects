package routers

import (
	"bookProject/internal/delivery"
	"bookProject/internal/middleware"
	"bookProject/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, bookService services.BookService) {
	r.POST("/auth/register", delivery.Register)
	r.POST("/auth/login", delivery.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/books", delivery.GetAllBooks(bookService))
		auth.GET("/books/:id", delivery.GetBookByID(bookService))

		admin := auth.Group("/")
		admin.Use(middleware.RequireRole("admin"))
		{
			admin.POST("/books", delivery.CreateBook(bookService))
			admin.PUT("/books/:id", delivery.UpdateBook(bookService))
			admin.DELETE("/books/:id", delivery.DeleteBook(bookService))
		}

		auth.GET("/auth/me", delivery.Me)
	}
}
