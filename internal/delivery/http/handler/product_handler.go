package handler

import (
	"errors"
	"net/http"
	"product_service/internal/delivery/http/dto/input"
	"product_service/internal/delivery/http/dto/request"
	"product_service/internal/delivery/http/dto/response"
	"product_service/internal/helper"
	"product_service/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase usecase.ProductUseCase
}

func NewProductHandler(u usecase.ProductUseCase) *ProductHandler {
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

    var resp []response.ProductResponse 

    for _, p := range product {
        resp = append(resp, response.ProductResponse{
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
    // ambil id dari param
    idParam := c.Param("id")
    id64, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
        return
    }

    product, err := h.usecase.GetProductByID(c.Request.Context(), id64)
    if err != nil {
        if errors.Is(err, usecase.ErrProductNotFound) {
            helper.ErrorResponse(c, http.StatusNotFound, "product not found")
            return
        }
        helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    resp := response.ProductResponse{
        ID:          int(product.ID),
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
	var req request.ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

    i := input.CreateProductInput{
        Name : req.Name,
        Description: req.Description,
        Price: req.Price,
        Image: req.Image,
    }

    product, err := h.usecase.CreateProduct(c.Request.Context(), i) 

    if err != nil {
        helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    resp := response.ProductResponse{
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

    var req request.ProductRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    i := input.UpdateProductInput{
        Name : req.Name,
        Description: req.Description,
        Price: req.Price,
        Image: req.Image,
    }

    product, err := h.usecase.UpdateProduct(c.Request.Context(), id64, i) 

    if err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    resp := response.ProductResponse{
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