package main

import (
	"go-ecommerce-api/config"
	"go-ecommerce-api/internal/handler"
	"go-ecommerce-api/internal/middleware"
	"go-ecommerce-api/internal/repository"
	"go-ecommerce-api/internal/routes"
	"go-ecommerce-api/internal/service"
	"go-ecommerce-api/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {

	logger.Init()
	config.LoadEnv()
	db := config.ConnectDB()
	config.RunMigrate(db)

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
	app.Use(middleware.RequestLogger())

	// setup routes
	routes.SetupRoutes(
		app,
		authHandler,
		productHandler,
		profileHandler,
		categoryHandler,
	)

	logger.Log.Info("env loaded",
		zap.String("port", config.GetEnv("APP_PORT")),
		zap.String("db_host", config.GetEnv("DB_HOST")),
	)
	app.Listen(":" + config.GetEnv("APP_PORT"))
}
