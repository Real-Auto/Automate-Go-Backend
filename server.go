package main

import (
	"Automate-Go-Backend/configs"
	"Automate-Go-Backend/routes"
	"Automate-Go-Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"net/http"


)

func main() {
    
	app := fiber.New()

	// Add the middleware to all endpoints
    app.Use(middleware.EnsureValidToken()(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// CORS Headers.
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
		}),
	))

	//run database
	configs.ConnectDB()

    //routes
    routes.Auth0Route(app)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    })

    app.Listen(":3000")
}
