package dto

type CreateProductInput struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Image string `json:"image"`
}

type UpdateProductInput struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Image string `json:"image"`
}