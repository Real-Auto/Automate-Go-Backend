package controllers

import (
	"Automate-Go-Backend/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DumbFunction(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "You are now Verified"})
}
