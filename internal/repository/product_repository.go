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

func (r *ProductRepositoryImpl) FindAllProduct(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product

	if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
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

func (r *ProductRepositoryImpl) Update(ctx context.Context, id int64, data map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&domain.Product{}).Where("id = ?", id).Updates(data).Error
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Product{}).Error
}