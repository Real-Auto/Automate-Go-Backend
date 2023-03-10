package routes

import (
	"Automate-Go-Backend/configs"
	"Automate-Go-Backend/controllers"
	"Automate-Go-Backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func auth0Route(app *fiber.App) {
    // Group endpoints that require authentication
	userPrivate := app.Group("/user-private", middleware.ValidateToken(configs.EnvGetUserScopes()))

    // private routes
    userPrivate.Post("/getUser", controllers.GetUser)
    userPrivate.Post("/changePassword", controllers.ChangePassword)
    userPrivate.Patch("/updateUser", controllers.UpdateUser)
    userPrivate.Delete("/deleteUser", controllers.DeleteUser)
    userPrivate.Get("/temp", controllers.Temp)
    userPrivate.Get("/temp2", controllers.Temp2)
    //implement logout endpoint

    // public routes
    app.Post("/signUp", controllers.SignUp)
    app.Post("/login", controllers.Login)

}
