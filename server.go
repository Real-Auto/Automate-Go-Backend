package main

import (
	"Automate-Go-Backend/configs"
	"Automate-Go-Backend/routes"

	"github.com/gofiber/fiber/v2"




)

func main() {
    
	app := fiber.New()




	// Add the middleware to all endpoints
	// app.Use("/user-private/*", fiberHandler)

	//run database
	configs.ConnectDB()

    //routes
    routes.Auth0Route(app)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    })

    app.Listen(":8080")
}
