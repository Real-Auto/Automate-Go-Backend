package routes

import (
	"Automate-Go-Backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func checkrRoutes(app *fiber.App) {
	//All routes related to users comes here
	app.Post("/GetUser", controllers.DumbFunction)
}
