package main

import (
	"clean-api/internal/delivery/http"
	"clean-api/internal/repository"
	"clean-api/internal/usecase"
	"clean-api/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize GORM database
	db, err := database.NewGormDB("./clean-api.db")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repositories
	userRepo := repository.NewGormUserRepository(db)
	productRepo := repository.NewGormProductRepository(db)
	orderRepo := repository.NewGormOrderRepository(db)

	// Initialize use cases
	userUseCase := usecase.NewUserUseCase(userRepo)
	productUseCase := usecase.NewProductUseCase(productRepo)
	orderUseCase := usecase.NewOrderUseCase(orderRepo)

	// Initialize Gin router
	r := gin.Default()

	// Initialize handlers
	http.NewUserHandler(r, userUseCase)
	http.NewProductHandler(r, productUseCase)
	http.NewOrderHandler(r, orderUseCase)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
