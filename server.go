package main

import (
	"Automate-Go-Backend/configs"
	"Automate-Go-Backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)
	routes.Auth0Route(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	app.Listen(":3000")
}
