package utils

import (
	"go-ecommerce-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET = []byte(config.GetEnv("secret"))

func GenerateJWT(userID uint, role string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SECRET)
}
