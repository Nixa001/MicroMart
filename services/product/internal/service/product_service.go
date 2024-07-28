package service

import (
	"errors"

	"github.com/Nixa001/micromart-services-product/internal/model"
	"github.com/Nixa001/micromart-services-product/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) ListProducts() ([]model.Product, error) {
	return s.repo.ListProducts()
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	// Vous pouvez ajouter ici une logique de validation si nécessaire
	return s.repo.CreateProduct(product)
}

func (s *ProductService) GetProduct(id string) (*model.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *ProductService) UpdateProduct(id string, product *model.Product) error {
	existingProduct, err := s.repo.GetProduct(id)
	if err != nil {
		return err
	}
	if existingProduct == nil {
		return errors.New("product not found")
	}

	// Mettez à jour les champs nécessaires
	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.Price = product.Price
	existingProduct.Category = product.Category

	return s.repo.UpdateProduct(existingProduct)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.repo.DeleteProduct(id)
}

func (s *ProductService) SearchProducts(query string) ([]model.Product, error) {
	return s.repo.SearchProducts(query)
}

func (s *ProductService) GetProductsByCategory(category string) ([]model.Product, error) {
	return s.repo.GetProductsByCategory(category)
}

func (s *ProductService) UpdateStock(id string, quantity int) error {
	product, err := s.repo.GetProduct(id)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	product.Stock = quantity
	return s.repo.UpdateProduct(product)
}
