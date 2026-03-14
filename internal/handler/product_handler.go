package handler

import (
	"go-ecommerce-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	Service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	search := c.Query("search")

	products, total, err := h.Service.GetProducts(page, limit, search)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"page":  page,
		"limit": limit,
		"total": total,
		"data":  products,
	})
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {

	type Request struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
		Stock int    `json:"stock"`
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.Service.CreateProduct(req.Name, req.Price, req.Stock)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "product created",
	})
}
