package routes

import "github.com/gofiber/fiber/v2"

import (
    "Automate-Go-Backend/controllers" 
)

func UserRoute(app *fiber.App) {
    //All routes related to users comes here
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Put("/user/:userId", controllers.EditAUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)
}