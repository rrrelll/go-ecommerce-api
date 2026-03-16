package config

import (
	"go-ecommerce-api/internal/model"

	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {

	db.AutoMigrate(

		&model.User{},
		&model.Product{},
		&model.Category{},
	)
}
