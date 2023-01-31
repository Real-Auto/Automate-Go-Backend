package main

import "github.com/gofiber/fiber/v2"

func routes(app fiber.App) *fiber.App {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/databasetest", func(c *fiber.Ctx) error {
		return c.SendString(Get())
	})
	return &app
}
