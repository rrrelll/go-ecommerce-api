package repository

import (
	"go-ecommerce-api/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {

	var products []model.Product

	err := r.DB.Find(&products).Error

	return products, err
}

func (r *ProductRepository) Create(product *model.Product) error {
	return r.DB.Create(product).Error
}

// ini contoh pagination
func (r *ProductRepository) GetWithPagination(page int, limit int, search string) ([]model.Product, int64, error) {

	var products []model.Product
	var total int64

	offset := (page - 1) * limit

	query := r.DB.Model(&model.Product{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query.Count(&total)

	err := query.Limit(limit).Offset(offset).Find(&products).Error

	return products, total, err
}
