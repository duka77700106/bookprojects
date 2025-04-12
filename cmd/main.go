package main

import (
	"bookProject/internal/config"
	"bookProject/internal/repository"
	"bookProject/internal/routers"
	"bookProject/internal/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := config.SetupDatabaseConnection()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	bookRepo := repository.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	r := gin.Default()
	routers.SetupRoutes(r, bookService)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
