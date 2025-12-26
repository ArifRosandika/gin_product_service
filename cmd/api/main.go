package main

import (
	"product_service/config"
	"product_service/internal/entity"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitConfig()
	r := gin.Default()

	db.AutoMigrate(&entity.Product{})

	r.Run(":8080")
}