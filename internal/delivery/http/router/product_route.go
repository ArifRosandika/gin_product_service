package router

import (
	"github.com/gin-gonic/gin"
	"product_service/internal/delivery/http/handler"
)
func NewProductRoute(r *gin.Engine, h *handler.ProductHandler) {
	r.GET("/products", h.GetAllProduct)
	r.GET("/products/:id", h.GetProductByID)
	r.POST("/products", h.CreateProduct)
	r.PUT("/products/:id", h.UpdateProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
}