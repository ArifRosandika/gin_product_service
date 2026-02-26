package domain

import (
	"context"
)

type ProductRepository interface {
	List(ctx context.Context, limit, offset int) ([]*Product, int, error)
	FindByID(ctx context.Context, id int64) (*Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id int64) error
}