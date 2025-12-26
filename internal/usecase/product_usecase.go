package usecase

import (
	"context"
	"product_service/internal/entity"
	"product_service/internal/repository"
	"product_service/internal/usecase/dto/input"
)

type ProductUseCaseImpl struct {
	repository *repository.ProductRepository
}

func NewProductUseCase(repo *repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		repository: repo,
	}
}

func (u *ProductUseCaseImpl) GetAllProduct(ctx context.Context) ([]*entity.Product, error) {
	return u.repository.FindAllProduct(ctx)
}

func (u *ProductUseCaseImpl) GetProductByID(ctx context.Context, id uint64) (*entity.Product, error) {
	return u.repository.FindByID(ctx, int64(id))
}

func (u *ProductUseCaseImpl) CreateProduct(ctx context.Context, input CreateProductInput) (*entity.Product, error) {
	return u.repository.Create(ctx, &entity.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Image:       input.Image,
	})
}

func (u *ProductUseCaseImpl) UpdateProduct(ctx context.Context, id uint64, input UpdateProductInput) (*entity.Product, error) {
	return u.repository.Update(ctx, *&entity.Product{
		ID:          int64(id),
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Image:       input.Image,
	})
}