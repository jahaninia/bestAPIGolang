package services

import (
    "context"
    "your_project/models"
    "your_project/repositories"
)

type ProductService interface {
    CreateProduct(ctx context.Context, product models.Product) (uint, error)
    GetProduct(ctx context.Context, id uint) (models.Product, error)
    DeleteProduct(ctx context.Context, id uint) error
    ListProducts(ctx context.Context) ([]models.Product, error)
}

type productService struct {
    repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
    return &productService{repo: repo}
}

func (s *productService) CreateProduct(ctx context.Context, product models.Product) (uint, error) {
    return s.repo.Create(ctx, product)
}

func (s *productService) GetProduct(ctx context.Context, id uint) (models.Product, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *productService) DeleteProduct(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *productService) ListProducts(ctx context.Context) ([]models.Product, error) {
    return s.repo.List(ctx)
}