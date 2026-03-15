package routes

import (
	"go-ecommerce-api/internal/handler"
	"go-ecommerce-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	authHandler *handler.AuthHandler,
	productHandler *handler.ProductHandler,
	profileHandler *handler.ProfileHandler,
	categoryHandler *handler.CategoryHandler,
) {

	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	app.Get("/profile",
		middleware.JWTProtected(),
		profileHandler.GetProfile,
	)

	app.Get("/products",
		middleware.JWTProtected(),
		productHandler.GetProducts,
	)

	app.Post("/products",
		middleware.JWTProtected(),
		middleware.RequireRole("seller"),
		productHandler.CreateProduct,
	)

	app.Put("/products/:id",
		middleware.JWTProtected(),
		middleware.RequireRole("seller"),
		productHandler.UpdateProduct,
	)

	app.Delete("/products/:id",
		middleware.JWTProtected(),
		middleware.RequireRole("seller"),
		productHandler.DeleteProduct,
	)

	app.Post("/categories",
		middleware.JWTProtected(),
		middleware.RequireRole("admin"),
		categoryHandler.CreateCategory,
	)

	app.Get("/categories",
		middleware.JWTProtected(),
		categoryHandler.GetCategories,
	)
}
