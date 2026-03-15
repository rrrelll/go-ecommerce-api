package service

import (
	"go-ecommerce-api/internal/model"
	"go-ecommerce-api/internal/repository"
	"go-ecommerce-api/pkg/utils"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) Register(name, email, password string) error {

	hash, _ := utils.HashPassword(password)

	user := model.User{
		Name:     name,
		Email:    email,
		Password: hash,
		Role:     "buyer",
	}

	return s.UserRepo.Create(&user)
}

func (s *AuthService) Login(email, password string) (string, error) {

	user, err := s.UserRepo.FindByEmail(email)

	if err != nil {
		return "", err
	}

	err = utils.CheckPassword(password, user.Password)

	if err != nil {
		return "", err
	}

	token, _ := utils.GenerateJWT(user.ID, user.Role)

	return token, nil
}

func (s *AuthService) GetProfile(userID uint) (*model.User, error) {
	return s.UserRepo.FindByID(userID)
}
