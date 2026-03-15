package handler

import (
	"strconv"

	"go-ecommerce-api/internal/dto"
	"go-ecommerce-api/internal/service"
	"go-ecommerce-api/internal/validation"
	"go-ecommerce-api/pkg/utils"

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
		return utils.Error(c, err.Error())
	}

	return utils.Success(c, "success", fiber.Map{
		"page":  page,
		"limit": limit,
		"total": total,
		"data":  products,
	})
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {

	var req dto.CreateProductRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, "invalid request body")
	}

	if err := validation.Validate.Struct(req); err != nil {
		return utils.Error(c, err.Error())
	}

	userID := uint(c.Locals("user_id").(float64))

	err := h.Service.CreateProduct(req, userID)

	if err != nil {
		return utils.Error(c, err.Error())
	}

	return utils.Success(c, "product created", nil)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	var req dto.UpdateProductRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, "invalid request")
	}

	if err := validation.Validate.Struct(req); err != nil {
		return utils.Error(c, err.Error())
	}

	userID := uint(c.Locals("user_id").(float64))

	err := h.Service.UpdateProduct(id, req, userID)

	if err != nil {
		return utils.Error(c, err.Error())
	}

	return utils.Success(c, "product updated", nil)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	userID := uint(c.Locals("user_id").(float64))

	err := h.Service.DeleteProduct(id, userID)

	if err != nil {
		return utils.Error(c, err.Error())
	}

	return utils.Success(c, "product deleted", nil)
}
