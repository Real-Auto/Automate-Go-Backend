package controllers

import (
    "Automate-Go-Backend/configs"
    "Automate-Go-Backend/models"
    "Automate-Go-Backend/responses"
    "encoding/json"
    "net/http"
    "bytes"
    "fmt"

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
	newUser := models.Auth0User{
		ClientID:    configs.EnvAuth0ClientID(),
		Connection:  configs.EnvAuth0Connection(),
		Email:       user.Email,
		Password:    user.Password,
		GivenName:   user.FirstName,
		FamilyName:  user.LastName,
		MetaData: map[string]string{
            "dateOfBirth": "1980-01-01",
            "gender":      "male",
        },
	}

    // Encode the user object into a JSON payload
	payload, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
    fmt.Println(user)

    // Create a HTTP post request
	r, err := http.NewRequest("POST", configs.EnvAuth0SignupEndpoint(), bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
	}

    // Set the HTTP request header.
    r.Header.Add("Content-Type", "application/json")
    client := &http.Client{}
    res, err := client.Do(r)
    fmt.Println(res.Body.Close())
    fmt.Println(err)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
    
    

    return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": res.Body.Close()}})


	
}

