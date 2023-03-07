package routes

import "github.com/gofiber/fiber/v2"

func RouteList(app *fiber.App) {
	auth0Route(app)
	checkrRoutes(app)
	faqRoutes(app)
}
