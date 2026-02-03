package handler

import (
	"net/http"
	"product_service/internal/delivery/http/dto"
	"product_service/internal/helper"
	"product_service/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase *usecase.ProductUseCaseImpl
}

func NewProductHandler(u *usecase.ProductUseCaseImpl) *ProductHandler {
	return &ProductHandler{
		usecase: u,
    }
}


func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	product, err := h.usecase.GetAllProduct(c. Request.Context())

    if err != nil {
        helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    var resp []dto.ProductResponse 

    for _, p := range product {
        resp = append(resp, dto.ProductResponse{
            ID : p.ID,
            Name : p.Name,
            Description: p.Description,
            Price: p.Price,
            Image: p.Image,
        })
    }

    helper.SuccessResponse(c, "success", gin.H{
        "products": resp,
    })
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
    idParam := c.Param("id")
    id64, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
        return
    }

    product, err := h.usecase.GetProductByID(c, id64)

    if err != nil {
        helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    resp := dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
	}

    helper.SuccessResponse(c, "success", gin.H{
        "product": resp,
    })
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

    i := dto.CreateProductInput{
        Name : req.Name,
        Description: req.Description,
        Price: req.Price,
        Image: req.Image,
    }

    product, err := h.usecase.CreateProduct(c.Request.Context(), &i) 

    if err != nil {
        helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    resp := dto.ProductResponse{
        ID:          int(product.ID),
        Name:        product.Name,
        Description: product.Description,
        Price:       product.Price,
        Image:       product.Image,
    }

    helper.SuccessResponse(c, "product create successfully", gin.H{
        "product": resp,
    })
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {

    idParam := c.Param("id")
    id64, err := strconv.ParseUint(idParam, 10, 64)

    if err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
        return
    }

    var req dto.ProductRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    in := dto.UpdateProductInput{
        Name : req.Name,
        Description: req.Description,
        Price: req.Price,
        Image: req.Image,
    }

    product, err := h.usecase.UpdateProduct(c.Request.Context(), id64, &in) 

    if err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    resp := dto.ProductResponse{
        ID : product.ID,
        Name : product.Name,
        Description: product.Description,
        Price: product.Price,
        Image: product.Image,
    }

    helper.SuccessResponse(c, "update product successfully", gin.H{
        "product": resp,
    })
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {

    idParam := c.Param("id")
    id64, err := strconv.ParseUint(idParam, 10, 64)

    if err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
        return
    }

    err = h.usecase.DeleteProduct(c.Request.Context(), id64)

    if err != nil {
        helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    helper.SuccessResponse(c, "delete product successfully", nil)
}