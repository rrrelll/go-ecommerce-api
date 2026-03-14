package middleware

import (
	"strings"

	"go-ecommerce-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {

	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Missing token",
			})
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return utils.SECRET, nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Locals("user_id", claims["user_id"])

		return c.Next()
	}
}
