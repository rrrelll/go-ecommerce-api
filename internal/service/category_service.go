package service

import (
	"go-ecommerce-api/internal/model"
	"go-ecommerce-api/internal/repository"
)

type CategoryService struct {
	Repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) CreateCategory(name string) error {

	category := model.Category{
		Name: name,
	}

	return s.Repo.Create(&category)
}

func (s *CategoryService) GetCategories() ([]model.Category, error) {
	return s.Repo.GetAll()
}
