package controllers

import (
	"Automate-Go-Backend/configs"
	"Automate-Go-Backend/models"
	"Automate-Go-Backend/responses"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"bytes"
	// "strings"
	// "reflect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

// swagger:operation GET /GetUser user GetUser
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
//	  description: User object
//	  schema:
//	    "$ref": "#/definitions/Auth0User"
func GetUser(c *fiber.Ctx) error {
	var user models.GetAuth0UserFieldsPayload
	validate := validator.New()
	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// Encode the user object into a JSON payload
	fmt.Println(user)
	payload, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error converting payload to JSON", Data: &fiber.Map{"data": err.Error()}})
	}

	// Create a HTTP post request
	r, err := http.NewRequest("GET", configs.EnvAuth0GetUserInfoEndpoint(), bytes.NewBuffer(payload))
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
	if jsErr := json.Unmarshal(body, &responseData); jsErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": "error unmarshelling response"}})
	}
	if err != nil {
		// handle error
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	fmt.Println(responseData)

	// parts := strings.Split(responseData["sub"].(string),"|")
	// if len(parts) != 2 {
	// 	fmt.Println("Unexpected format")
	// 	return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": "auth0 auth0Response[\"sub\"] could not be split by |"}})
	// }

	// var userResponse models.Auth0User

	// objId, _ := primitive.ObjectIDFromHex(parts[1])

	// mongErr := userCollection.FindOne(c.Context() ,bson.M{"_id": objId}).Decode(&userResponse)
	// if mongErr != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	// }

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": responseData}})

}

// swagger:operation POST /signUp user signUp
//
// # Sign up endpoint
//
// This endpoint returns a confirmation message.
//
// ---
// produces:
// - application/json
// parameters:
//
// responses:
//
//	200:
//	  description: Success message
func SignUp(c *fiber.Ctx) error {
	var user models.SignUpPayload
	validate := validator.New()
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
		ClientId:   configs.EnvAuth0ClientId(),
		Connection: configs.EnvAuth0Connection(),
		Email:      user.Email,
		Password:   user.Password,
		GivenName:  user.FirstName,
		FamilyName: user.LastName,
		Name:       user.Name,
		MetaData: models.UserMetaData{
			Services:     user.Services,
			DateOfBirth:  user.DateOfBirth,
			PhotoFileUrl: user.PhotoFileUrl,
			Phone:        user.Phone,
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
	if jsErr := json.Unmarshal(body, &responseData); jsErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": "error unmarshelling response"}})
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

// swagger:operation POST /login user Login
//
// # Login in endpoint
//
// This endpoint returns a confirmation message.
//
// ---
// produces:
// - application/json
// parameters:
//
// responses:
//
//	200:
//	  description: Success message
func Login(c *fiber.Ctx) error {
	var user models.LoginPayload
	validate := validator.New()
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
		GrantType:    "password",
		ClientId:     configs.EnvAuth0ClientId(),
		ClientSecret: configs.EnvAuth0ClientSecret(),
		// Audience:   configs.EnvAuth0ApiAudience(),
		Scope:    "openid profile email",
		Email:    user.Email,
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
	if jsErr := json.Unmarshal(body, &responseData); jsErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": "error unmarshelling response"}})
	}
	if err != nil {
		// handle error
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	fmt.Println(responseData)

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": responseData}})

}

// swagger:operation POST /changePassword user
//
// # changePassword endpoint
//
// This endpoint returns a confirmation message.
//
// ---
// produces:
// - application/json
// parameters:
//
// responses:
//
//	200:
//	  description: Success message
func ChangePassword(c *fiber.Ctx) error {
	var user models.ChangePasswordPayload
	validate := validator.New()
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
		ClientId:   configs.EnvAuth0ClientId(),
		Email:      user.Email,
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
	if jsErr := json.Unmarshal(body, &responseData); jsErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": "error unmarshelling response"}})
	}
	if err != nil {
		// handle error
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": responseData}})

}

// swagger:operation POST /updateUser user
//
// # update User endpoint
//
// This endpoint returns a confirmation message.
//
// ---
// produces:
// - application/json
// parameters:
//
// responses:
//
//	200:
//	  description: Success message
func UpdateUser(c *fiber.Ctx) error {

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": "data"}})
}

// swagger:operation POST /deleteUser user
//
// # delete User endpoint
//
// This endpoint returns a confirmation message.
//
// ---
// produces:
// - application/json
// parameters:
//
// responses:
//
//	200:
//	  description: Success message
func DeleteUser(c *fiber.Ctx) error {
	return nil
	// userId := c.Params("user_id")

	// if userId == "" {
	// 	// handle error
	// 	return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": "User ID parameter is missing"}})
	// }

	// url := fmt.Sprintf("https://dev-o3njyhd54d52dwd8.us.auth0.com/api/v2/users/%s", userId)

	// req, err := http.NewRequest("DELETE", url), nil)
	// if err != nil {
	// 	// handle error
	// 	return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	// }

	// client := &http.Client{}
	// res, err := client.Do(req)
	// if err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	// }
	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// var responseData map[string]interface{}
	// if jsErr := json.Unmarshal(body, &responseData); jsErr != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": "error unmarshelling response"}})
	// }
	// if err != nil {
	// 	// handle error
	// 	return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	// }

	// return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": responseData}})
}
