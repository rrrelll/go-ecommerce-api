package service

import (
	"errors"
	"go-ecommerce-api/internal/dto"
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

func (s *ProductService) CreateProduct(req dto.CreateProductRequest, userID uint) error {

	product := model.Product{
		Name:       req.Name,
		Price:      req.Price,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
		UserID:     userID,
	}

	return s.ProductRepo.Create(&product)
}

func (s *ProductService) UpdateProduct(id int, req dto.UpdateProductRequest, UserID uint) error {
	product, err := s.ProductRepo.FindByID(id)

	if err != nil {
		return err
	}

	if product.UserID != UserID {

		return errors.New("not product owner")
	}

	product.Name = req.Name
	product.Price = req.Price
	product.Stock = req.Stock
	product.CategoryID = req.CategoryID

	return s.ProductRepo.Update(product)
}

func (s *ProductService) DeleteProduct(id int, UserID uint) error {

	product, err := s.ProductRepo.FindByID(id)

	if err != nil {
		return err
	}

	if product.UserID != UserID {
		return errors.New("not product owner")
	}

	return s.ProductRepo.Delete(id)
}
