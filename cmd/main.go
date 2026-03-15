package main

import (
	"go-ecommerce-api/config"
	"go-ecommerce-api/internal/handler"
	"go-ecommerce-api/internal/repository"
	"go-ecommerce-api/internal/routes"
	"go-ecommerce-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db := config.ConnectDB()

	// repository
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	// service
	authService := service.NewAuthService(userRepo)
	productService := service.NewProductService(productRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	// handler
	authHandler := handler.NewAuthHandler(authService)
	productHandler := handler.NewProductHandler(productService)
	profileHandler := handler.NewProfileHandler(authService)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	app := fiber.New()

	// setup routes
	routes.SetupRoutes(
		app,
		authHandler,
		productHandler,
		profileHandler,
		categoryHandler,
	)

	app.Listen(":3000")
}
