package handler

import (
	"go-ecommerce-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {

	type Request struct {
		Name     string
		Email    string
		Password string
	}

	var req Request

	c.BodyParser(&req)

	err := h.Service.Register(req.Name, req.Email, req.Password)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.JSON("register success")
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {

	type Request struct {
		Email    string
		Password string
	}

	var req Request

	c.BodyParser(&req)

	token, err := h.Service.Login(req.Email, req.Password)

	if err != nil {
		return c.Status(401).JSON("login failed")
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
