package middleware

import (
	"go-ecommerce-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func RequireRole(role string) fiber.Handler {

	return func(c *fiber.Ctx) error {

		userRole := c.Locals("role").(string)

		if userRole != role {
			return utils.Error(c, "forbidden")
		}

		return c.Next()
	}
}
