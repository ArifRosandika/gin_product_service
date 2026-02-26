package repository

import (
	"context"
	"product_service/internal/domain"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (r *ProductRepositoryImpl) List(ctx context.Context, limit, offset int) ([]*domain.Product, int, error) {

	var (
		product []*domain.Product
		total int64
	)

	if err := r.db.WithContext(ctx).Model(&domain.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&product).Error; err != nil {
		return nil, 0, err
	}

	return product, int(total), nil
}

func (r *ProductRepositoryImpl) FindByID(ctx context.Context, id int64) (*domain.Product, error) {
	var product domain.Product

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepositoryImpl) Create(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).Create(&product).Error
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).Save(&product).Error
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Product{}).Error
}