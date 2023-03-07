package controllers

import (
	"Automate-Go-Backend/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// swagger:operation GET /Check check
//
// Get a user.
//
// This endpoint returns a user object.
//
// ---
// produces:
// - application/json
// parameters:
//
// responses:
//
//	200:
//	  description: message
func DumbFunction(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "You are now Verified"})
}
