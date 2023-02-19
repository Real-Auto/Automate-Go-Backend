package routes

import "github.com/gofiber/fiber/v2"

func RouteList(app *fiber.App) {
	userRoute(app)
	auth0Route(app)
	checkrRoutes(app)
}
