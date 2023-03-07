package controllers

import (
	"Automate-Go-Backend/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

// swagger:operation GET /GetFaq FAQ
//
// # FAQs
//
// This endpoint a FAQ JSON.
//
// ---
// produces:
// - application/json
// parameters:
//
// responses:
//
//	200:
//	  description: User object
func GetFaq(c *fiber.Ctx) error {
	file, err := os.Open("FAQ.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err}})
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err}})
	}

	var FAQ map[string]interface{}
	err2 := json.Unmarshal(data, &FAQ)
	if err2 != nil {
		// Handle error
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": FAQ}})
}
