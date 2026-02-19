package dto

type ProductRequest struct {
	Name string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"required,min=3"`
	Price float64 `json:"price" validate:"gt=0"`
	Image string `json:"image" validate:"required,url"`
}