package repository

import (
	"go-ecommerce-api/internal/model"
	"go-ecommerce-api/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	var user model.User

	err := r.DB.Where("email = ?", email).First(&user).Error

	return &user, err
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {

	var user model.User

	err := r.DB.First(&user, id).Error

	if err != nil {

		//Logging FindByID
		logger.Log.Error("Database Error!",
			zap.Error(err),
		)

		return nil, err
	}

	return &user, nil
}
