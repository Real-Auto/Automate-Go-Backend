package routes

import (
	"Automate-Go-Backend/controllers"
	"Automate-Go-Backend/middleware"
	"Automate-Go-Backend/configs"

	"github.com/gofiber/fiber/v2"
)

func faqRoutes(app *fiber.App) {
	// Group endpoints that require authentication
	userPrivate := app.Group("/user-private", middleware.ValidateToken(configs.EnvGetUserScopes()))

	// All routes related to users comes here
	userPrivate.Get("/faq", controllers.GetFaq)
}
