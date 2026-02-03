package usecase

import (
	"context"
	"product_service/internal/delivery/http/dto"
	"product_service/internal/domain"
)

type ProductUseCaseImpl struct {
	repository domain.ProductRepository
}

func NewProductUseCase(repository domain.ProductRepository) *ProductUseCaseImpl {
	return &ProductUseCaseImpl{
		repository: repository,
	}
}

func (u *ProductUseCaseImpl) GetAllProduct(ctx context.Context) ([]*domain.Product, error) {
	return u.repository.FindAllProduct(ctx)
}

func (u *ProductUseCaseImpl) GetProductByID(ctx context.Context, id uint64) (*domain.Product, error) {
	return u.repository.FindByID(ctx, int64(id))
}

func (u *ProductUseCaseImpl) CreateProduct(ctx context.Context, input *dto.CreateProductInput) (*domain.Product, error) {
	product := &domain.Product {
		Name : input.Name,
		Description : input.Description,
		Price : input.Price,
		Image : input.Image,
	}

	if err := u.repository.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (u *ProductUseCaseImpl) UpdateProduct(ctx context.Context, id uint64, input *dto.UpdateProductInput) (*domain.Product, error) {
 if err := u.repository.Update(ctx, int64(id), map[string]interface{}{

		"name":        input.Name,
		"description": input.Description,
		"price":       input.Price,
		"image":       input.Image,
	},
); err != nil {
		return nil, err
	}

	return u.repository.FindByID(ctx, int64(id))
}

func (u *ProductUseCaseImpl) DeleteProduct(ctx context.Context, id uint64) error {
	return u.repository.Delete(ctx, int64(id))
}