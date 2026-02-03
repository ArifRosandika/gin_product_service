package main

import (
	"product_service/config"
	"product_service/internal/domain"
	"product_service/internal/usecase"
	"product_service/internal/repository"
	"product_service/internal/delivery/http/router"
	"product_service/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitConfig()
	r := gin.Default()

	db.AutoMigrate(&domain.Product{})

	productRepository := repository.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)

	router.NewProductRoute(r, productHandler)

	r.Run(":8080")
}