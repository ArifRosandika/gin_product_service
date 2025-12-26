package usecase

import (
	"product_service/internal/entity"
	"context"
)

type ProductUseCase interface {
	GetAllProduct(ctx context.Context) ([]*entity.Product, error)
	GetProductByID(ctx context.Context, id uint64) (*entity.Product, error)
	CreateProduct(ctx context.Context, input CreateProductInput) (*entity.Product, error)
	UpdateProduct(ctx context.Context, id uint64, input UpdateProductInput) (*entity.Product, error)
	DeleteProduct(ctx context.Context, id uint64) error
}