package repository

import (
	"context"
	"product_service/internal/entity"
)

type ProductRepository interface {
	FindAllProduct(ctx context.Context) ([]*entity.Product, error) 
	FindByID(ctx context.Context, id int64) (*entity.Product, error)
	Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
	Update(ctx context.Context, product *entity.Product) (*entity.Product, error)
	Delete(ctx context.Context, id int64) error
}