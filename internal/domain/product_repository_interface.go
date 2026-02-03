package domain

import (
	"context"
)

type ProductRepository interface {
	FindAllProduct(ctx context.Context) ([]*Product, error)
	FindByID(ctx context.Context, id int64) (*Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, id int64, data map[string]interface{}) error 
	Delete(ctx context.Context, id int64) error
}