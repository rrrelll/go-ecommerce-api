package main

import (
	"go-ecommerce-api/config"
	"go-ecommerce-api/internal/handler"
	"go-ecommerce-api/internal/middleware"
	"go-ecommerce-api/internal/repository"
	"go-ecommerce-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db := config.ConnectDB()

	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	authService := service.NewAuthService(userRepo)
	productService := service.NewProductService(productRepo)

	authHandler := handler.NewAuthHandler(authService)
	productHandler := handler.NewProductHandler(productService)
	profileHandler := handler.NewProfileHandler(authService)

	app := fiber.New()

	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	app.Get("/products", middleware.JWTProtected(), productHandler.GetProducts)
	app.Post("/products", middleware.JWTProtected(), productHandler.CreateProduct)

	app.Get("/profile", middleware.JWTProtected(), profileHandler.GetProfile)

	app.Listen(":3000")
}
