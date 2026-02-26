package handler

import (
	"net/http"
	"product_service/internal/delivery/http/dto"
	"product_service/internal/domain"
	"product_service/internal/helper"
	"product_service/internal/usecase"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	usecase *usecase.ProductUseCaseImpl
}

func NewProductHandler(u *usecase.ProductUseCaseImpl) *ProductHandler {
	return &ProductHandler{
		usecase: u,
    }
}

func (h *ProductHandler) List(c *gin.Context) {
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

    if limit > 50 {
        limit = 50
    }

    if page < 1 {
        page = 1
    }

    offset := (page - 1) * limit
    

    items, total, err := h.usecase.List(c, limit, offset)

    if err != nil {
        helper.ErrorResponse(c, 500, gin.H{
            "message" : err.Error(),
        })
    }

    responses := make([]dto.ProductResponse, 0, len(items))
    for _, p := range items {
        responses = append(responses, dto.ProductResponse {
            ID:          p.ID,
            Name:        p.Name,
            Description: p.Description,
            Price:       p.Price,
            Image:       p.Image,
        })
    }

    helper.SuccessResponse(c, "products fetched", gin.H{
        "items": responses,
        "meta": gin.H{
            "page": page,
            "limit": limit,
            "total": total,
        },
    })
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
    idParam := c.Param("id")
    id64, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, gin.H{
            "message": "invalid id",
        })
        return
    }

    product, err := h.usecase.GetProductByID(c, id64)

    if err != nil {
        helper.ErrorResponse(c, http.StatusInternalServerError, gin.H{
            "message": err.Error(),
        })
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
		helper.ErrorResponse(c, http.StatusBadRequest, gin.H{
            "message": err.Error(),
        })
		return
	}

    validators := validator.New()

    if err := validators.Struct(req); err != nil {
        errors := make(map[string]string) 

        for _, e := range err.(validator.ValidationErrors) {
            field := strings.ToLower(e.Field())

            switch e.Tag() {
            case "required":
                errors[field] = field + " is required"
            case "min":
                errors[field] = field + " must be at least " + e.Param()
            case "gt":
                errors[field] = field + " must be greater than " + e.Param()
            case "url":
                errors[field] = field + " must be a valid URL"
            default:
                errors[field] = e.Error()
            }
        }

        helper.ErrorResponse(c, 400, gin.H{
            "message": "validation failed",
            "errors":  errors,
        })

        return
    }

    p := &domain.Product{
        Name :       req.Name,
        Description: req.Description,
        Price:       req.Price,
        Image:       req.Image,
    }

    product, err := h.usecase.CreateProduct(c.Request.Context(), p) 

    if err != nil {
        helper.ErrorResponse(c, http.StatusInternalServerError, gin.H{
            "message": err.Error(),
        })
        return
    }

    resp := dto.ProductResponse{
        ID:          product.ID,
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
        helper.ErrorResponse(c, http.StatusBadRequest, gin.H{
            "message": "invalid id",
        })
        return  
    }

    var req dto.ProductRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, gin.H{
            "message": err.Error(),
        })
        return
    }

    validators := validator.New()

    if err := validators.Struct(req); err != nil {
        errors := make(map[string]string) 

        for _, e := range err.(validator.ValidationErrors) {
            field := strings.ToLower(e.Field())

            switch e.Tag() {
            case "required":
                errors[field] = field + " is required"
            case "min":
                errors[field] = field + " must be at least " + e.Param()
            case "gt":
                errors[field] = field + " must be greater than " + e.Param()
            case "url":
                errors[field] = field + " must be a valid URL"
            default:
                errors[field] = err.Error()
            }
        }

        helper.ErrorResponse(c, 400, gin.H{
            "message": "validation failed",
            "errors":  errors,
        })

        return
    }

    product, err := h.usecase.UpdateProduct(c.Request.Context(), id64, &dto.ProductRequest{
        Name :       req.Name,
        Description: req.Description,
        Price:       req.Price,
        Image:       req.Image,
    })

    if err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, gin.H{
            "message": err.Error(),
        })
        return
    }

    resp := dto.ProductResponse{
        ID :         product.ID,
        Name :       product.Name,
        Description: product.Description,
        Price:       product.Price,
        Image:       product.Image,
    }

    helper.SuccessResponse(c, "update product successfully", gin.H{
        "product": resp,
    })
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {

    idParam := c.Param("id")
    id64, err := strconv.ParseUint(idParam, 10, 64)

    if err != nil {
        helper.ErrorResponse(c, http.StatusBadRequest, gin.H{
            "message": "invalid id",
        })
        return
    }

    err = h.usecase.DeleteProduct(c.Request.Context(), id64)

    if err != nil {
        helper.ErrorResponse(c, http.StatusInternalServerError, gin.H{
            "message": err.Error(),
        })
        return
    }

    helper.SuccessResponse(c, "delete product successfully", nil)
}