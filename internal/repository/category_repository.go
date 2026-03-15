package repository

import (
	"go-ecommerce-api/internal/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (r *CategoryRepository) Create(category *model.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) GetAll() ([]model.Category, error) {

	var categories []model.Category

	err := r.DB.Find(&categories).Error

	return categories, err
}
