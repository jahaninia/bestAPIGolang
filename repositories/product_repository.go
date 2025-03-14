package repositories

import (
	"context"
	"your_project/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, product models.Product) (uint, error)
	FindByID(ctx context.Context, id uint) (models.Product, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]models.Product, error)
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{DB: db}
}

func (r *productRepository) Create(ctx context.Context, product models.Product) (uint, error) {
	result := r.DB.WithContext(ctx).Create(&product)
	return product.ID, result.Error
}

func (r *productRepository) FindByID(ctx context.Context, id uint) (models.Product, error) {
	var product models.Product
	result := r.DB.WithContext(ctx).First(&product, id)
	return product, result.Error
}

func (r *productRepository) Delete(ctx context.Context, id uint) error {
	result := r.DB.WithContext(ctx).Delete(&models.Product{}, id)
	return result.Error
}

func (r *productRepository) List(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	result := r.DB.WithContext(ctx).Find(&products)
	return products, result.Error
}
