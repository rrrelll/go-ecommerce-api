package handler

import (
	"go-ecommerce-api/internal/service"
	"go-ecommerce-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	Service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {

	type Request struct {
		Name string `json:"name"`
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, "invalid request")
	}

	err := h.Service.CreateCategory(req.Name)

	if err != nil {
		return utils.Error(c, err.Error())
	}

	return utils.Success(c, "category created", nil)
}

func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {

	categories, err := h.Service.GetCategories()

	if err != nil {
		return utils.Error(c, err.Error())
	}

	return utils.Success(c, "success", categories)
}
