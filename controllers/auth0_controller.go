package controllers

import (
	"Automate-Go-Backend/configs"
	"Automate-Go-Backend/models"
	"Automate-Go-Backend/responses"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	var user models.User

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// create user for model
	newUser := models.Auth0User {
		ClientId:   configs.EnvAuth0ClientId(),
		Connection: configs.EnvAuth0Connection(),
		Email:      user.Email,
		Password:   user.Password,
		GivenName:  user.FirstName,
		FamilyName: user.LastName,
		MetaData:  models.UserMetaData{
			Services:     user.Services,
			DateOfBirth: user.DateOfBirth,
			PhotoFileUrl:  user.PhotoFileUrl,
			Phone: user.Phone,
		},
	}

	// Encode the user object into a JSON payload
	payload, err := json.Marshal(newUser)
	fmt.Println(payload)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error converting payload to JSON", Data: &fiber.Map{"data": err.Error()}})
	}

	// Create a HTTP post request
	r, err := http.NewRequest("POST", configs.EnvAuth0SignupEndpoint(), bytes.NewBuffer(payload))
	//fmt.Println(r)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// Set the HTTP request header.
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		fmt.Println("Error unmarshaling response:", err)
	}
	if err != nil {
		// handle error
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	fmt.Println(bytes.NewBuffer(body))

	//if user already exists return appropriate error response
	if responseData["code"] == "user_exists" {
		return c.Status(http.StatusConflict).JSON(responses.UserResponse{Status: http.StatusConflict, Message: "error", Data: &fiber.Map{"data": "user already exists"}})
	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success User Created", Data: &fiber.Map{"data": responseData}})

}

func Login(c *fiber.Ctx) error {
	var user models.LoginPayload

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// create user for model
	newUser := models.Auth0UserLogin{
		GrantType:   "password",
		ClientId: configs.EnvAuth0ClientId(),
		ClientSecret:   configs.EnvAuth0ClientSecret(),
		// Audience:   configs.EnvAuth0ApiAudience(),
		Email:  user.Email,
		Password: user.Password,
	}

	// Encode the user object into a JSON payload
	payload, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error converting payload to JSON", Data: &fiber.Map{"data": err.Error()}})
	}

	// Create a HTTP post request
	r, err := http.NewRequest("POST", configs.EnvAuth0LoginEndpoint(), bytes.NewBuffer(payload))
	//fmt.Println(r)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// Set the HTTP request header.
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		fmt.Println("Error unmarshaling response:", err)
	}
	if err != nil {
		// handle error
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	fmt.Println(responseData)

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": responseData}})

}

func ChangePassword(c *fiber.Ctx) error {
	var user models.ChangePasswordPayload

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// create user for model
	newUser := models.Auth0UserChangePassword{
        ClientId: configs.EnvAuth0ClientId(),
        Email: user.Email,
        Connection: configs.EnvAuth0Connection(),
	}

	// Encode the user object into a JSON payload
	payload, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error converting payload to JSON", Data: &fiber.Map{"data": err.Error()}})
	}

	// Create a HTTP post request
	r, err := http.NewRequest("POST", configs.EnvAuth0ChangePasswordEndpoint(), bytes.NewBuffer(payload))
	//fmt.Println(r)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// Set the HTTP request header.
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		fmt.Println("Error unmarshaling response:", err)
	}
	if err != nil {
		// handle error
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}


	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": responseData}})

}

