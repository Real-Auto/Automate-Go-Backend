//	Title: AutoMate API
//	BasePath: /v1
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//
//	SecurityDefinitions:
//	api_key:
//	     type: apiKey
//	     name: KEY
//	     in: header
//	oauth2:
//	    type: oauth2
//	    authorizationUrl: /oauth2/auth
//	    tokenUrl: /oauth2/token
//	    in: header
//	    scopes:
//	      bar: foo
//	    flow: accessCode
//
//	Extensions:
//	x-meta-value: value
//	x-meta-array:
//	  - value1
//	  - value2
//	x-meta-array-obj:
//	  - name: obj
//	    value: field
//
// swagger:meta
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
	routes.RouteList(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

    app.Listen(":8080")
}
