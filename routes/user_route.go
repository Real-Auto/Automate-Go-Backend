package routes

import (
	"Automate-Go-Backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func userRoute(app *fiber.App) {
	//All routes related to users comes here
	// app.Post("/user", controllers.CreateUser)
	// app.Patch("/user/:userId", controllers.EditProfileInformation)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)
}
