package handler

import (
	"go-ecommerce-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	Service *service.AuthService
}

func NewProfileHandler(service *service.AuthService) *ProfileHandler {
	return &ProfileHandler{service}
}

func (h *ProfileHandler) GetProfile(c *fiber.Ctx) error {

	// user_id biasanya disimpan di context oleh middleware JWT
	userID := c.Locals("user_id")

	if userID == nil {
		return c.Status(401).JSON("unauthorized")
	}

	id := uint(userID.(float64))

	user, err := h.Service.GetProfile(id)

	if err != nil {
		return c.Status(404).JSON("user not found")
	}

	return c.JSON(user)
}
