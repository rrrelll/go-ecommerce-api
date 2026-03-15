package dto

type CreateProductRequest struct {
	Name       string `json:"name" validate:"required,min=3"`
	Price      int    `json:"price" validate:"required,gt=0"`
	Stock      int    `json:"stock" validate:"required,gte=0"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
	Name       string `json:"name" validate:"required,min=3"`
	Price      int    `json:"price" validate:"required,gt=0"`
	Stock      int    `json:"stock" validate:"required,gte=0"`
	CategoryID uint   `json:"category_id" validate:"required"`
}
