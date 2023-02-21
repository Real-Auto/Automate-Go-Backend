package routes

import "github.com/gofiber/fiber/v2"

import (
    "Automate-Go-Backend/controllers" 
)

func Auth0Route(app *fiber.App) {
    //All routes related to users comes here
    app.Post("/GetUser", controllers.GetUser)
	app.Post("/signUp", controllers.SignUp)
    app.Post("/login", controllers.Login)
    app.Post("/changePassword", controllers.ChangePassword)
    app.Patch("/updateUser", controllers.UpdateUser)
    app.Delete("/deleteUser", controllers.DeleteUser)
    //implement logout endpoint


}