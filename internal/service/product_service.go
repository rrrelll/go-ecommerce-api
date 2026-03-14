package service

import (
	"go-ecommerce-api/internal/model"
	"go-ecommerce-api/internal/repository"
)

type ProductService struct {
	ProductRepo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo}
}

// ini contoh pagination
func (s *ProductService) GetProducts(page int, limit int, search string) ([]model.Product, int64, error) {

	return s.ProductRepo.GetWithPagination(page, limit, search)
}

func (s *ProductService) CreateProduct(name string, price int, stock int) error {

	product := model.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}

	return s.ProductRepo.Create(&product)
}
