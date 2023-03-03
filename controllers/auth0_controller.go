package controllers

import (
	"Automate-Go-Backend/configs"
	"Automate-Go-Backend/models"
	"Automate-Go-Backend/responses"
	"Automate-Go-Backend/middleware"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/go-playground/validator/v10"
	"bytes"
	//"strings"
	"reflect"
	//"gopkg.in/auth0.v5"
	// "gopkg.in/auth0.v5/management"
	// "github.com/auth0/go-auth0"
	"github.com/auth0/go-auth0/management"
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/gofiber/fiber/v2"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

type MyStruct struct {
	Field1 string   `json:"field1"`
	Field2 string `json:"field2"`
}

func Temp(c *fiber.Ctx) error {
	fmt.Println(middleware.Auth0TokenVar)
	res, err := middleware.GetManagementApiToken()
	if err != nil {
		fmt.Println(reflect.TypeOf(err))
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": res}})
}

func Temp2(c *fiber.Ctx) error {
	fmt.Println(middleware.Auth0TokenVar.ExpiresIn)
	fmt.Println(middleware.Expiration)
	// convert to time object
	expiryTime := time.Unix(middleware.Expiration, 0)

	// check if the expiry time has passed
	if time.Now().After(expiryTime) {
		// access token has expired, take appropriate action
		fmt.Println("expired.." + " time now is: " + time.Now().Format("2006-01-02 15:04:05") + "\nexpiration time is: " + expiryTime.Format("2006-01-02 15:04:05"))
	} else {
		// access token is still valid
		fmt.Println("NOT expired.." + " time now is: " + time.Now().Format("2006-01-02 15:04:05") + "\nexpiration time is: " + expiryTime.Format("2006-01-02 15:04:05"))
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "token printed to console"}})
}

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
	fmt.Println(string(body))
	if err != nil {
		// handle error
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	
	if (string(body) == "Unauthorized") {
		return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{Status: http.StatusUnauthorized, Message: "error", Data: &fiber.Map{"data": "Unauthorized"}})
	} 

	var responseData models.GetAuth0UserResponse
	if jsErr := json.Unmarshal(body, &responseData); jsErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": "error unmarshelling response"}})
	}
	
	fmt.Println(responseData)

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
			Language: 	  user.Language,
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
		GrantType:   "password",
		ClientId: configs.EnvAuth0ClientId(),
		ClientSecret:   configs.EnvAuth0ClientSecret(),
		Audience:   configs.EnvAuth0ApiAudience(),
		Scope: "openid profile email",
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
	var user models.UpdateAuth0UserPayload;

	// validate request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	//split user ID
	user_id_by_itself := strings.Split(user.UserId,"|")[1]
	fmt.Println(user_id_by_itself)

	// check if user exists on mongo end
	var mongo_user models.Auth0User

	objId, _ := primitive.ObjectIDFromHex(user_id_by_itself)

	err := userCollection.FindOne(c.Context(), bson.M{"_id": objId}).Decode(&mongo_user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": "user does not exist"}})
	}

	// get access_token for user
	access_token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IllxSnViYlM5U290UzNVdUtYVk9hbyJ9.eyJpc3MiOiJodHRwczovL2Rldi1vM25qeWhkNTRkNTJkd2Q4LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJjOTkwYmw1T2tRcjhJNjdYNHNuMDNudzNrREtSUEoyYkBjbGllbnRzIiwiYXVkIjoiaHR0cHM6Ly9kZXYtbzNuanloZDU0ZDUyZHdkOC51cy5hdXRoMC5jb20vYXBpL3YyLyIsImlhdCI6MTY3NzYxNDQzOSwiZXhwIjoxNjc3NzAwODM5LCJhenAiOiJjOTkwYmw1T2tRcjhJNjdYNHNuMDNudzNrREtSUEoyYiIsInNjb3BlIjoicmVhZDpjbGllbnRfZ3JhbnRzIGNyZWF0ZTpjbGllbnRfZ3JhbnRzIGRlbGV0ZTpjbGllbnRfZ3JhbnRzIHVwZGF0ZTpjbGllbnRfZ3JhbnRzIHJlYWQ6dXNlcnMgdXBkYXRlOnVzZXJzIGRlbGV0ZTp1c2VycyBjcmVhdGU6dXNlcnMgcmVhZDp1c2Vyc19hcHBfbWV0YWRhdGEgdXBkYXRlOnVzZXJzX2FwcF9tZXRhZGF0YSBkZWxldGU6dXNlcnNfYXBwX21ldGFkYXRhIGNyZWF0ZTp1c2Vyc19hcHBfbWV0YWRhdGEgcmVhZDp1c2VyX2N1c3RvbV9ibG9ja3MgY3JlYXRlOnVzZXJfY3VzdG9tX2Jsb2NrcyBkZWxldGU6dXNlcl9jdXN0b21fYmxvY2tzIGNyZWF0ZTp1c2VyX3RpY2tldHMgcmVhZDpjbGllbnRzIHVwZGF0ZTpjbGllbnRzIGRlbGV0ZTpjbGllbnRzIGNyZWF0ZTpjbGllbnRzIHJlYWQ6Y2xpZW50X2tleXMgdXBkYXRlOmNsaWVudF9rZXlzIGRlbGV0ZTpjbGllbnRfa2V5cyBjcmVhdGU6Y2xpZW50X2tleXMgcmVhZDpjb25uZWN0aW9ucyB1cGRhdGU6Y29ubmVjdGlvbnMgZGVsZXRlOmNvbm5lY3Rpb25zIGNyZWF0ZTpjb25uZWN0aW9ucyByZWFkOnJlc291cmNlX3NlcnZlcnMgdXBkYXRlOnJlc291cmNlX3NlcnZlcnMgZGVsZXRlOnJlc291cmNlX3NlcnZlcnMgY3JlYXRlOnJlc291cmNlX3NlcnZlcnMgcmVhZDpkZXZpY2VfY3JlZGVudGlhbHMgdXBkYXRlOmRldmljZV9jcmVkZW50aWFscyBkZWxldGU6ZGV2aWNlX2NyZWRlbnRpYWxzIGNyZWF0ZTpkZXZpY2VfY3JlZGVudGlhbHMgcmVhZDpydWxlcyB1cGRhdGU6cnVsZXMgZGVsZXRlOnJ1bGVzIGNyZWF0ZTpydWxlcyByZWFkOnJ1bGVzX2NvbmZpZ3MgdXBkYXRlOnJ1bGVzX2NvbmZpZ3MgZGVsZXRlOnJ1bGVzX2NvbmZpZ3MgcmVhZDpob29rcyB1cGRhdGU6aG9va3MgZGVsZXRlOmhvb2tzIGNyZWF0ZTpob29rcyByZWFkOmFjdGlvbnMgdXBkYXRlOmFjdGlvbnMgZGVsZXRlOmFjdGlvbnMgY3JlYXRlOmFjdGlvbnMgcmVhZDplbWFpbF9wcm92aWRlciB1cGRhdGU6ZW1haWxfcHJvdmlkZXIgZGVsZXRlOmVtYWlsX3Byb3ZpZGVyIGNyZWF0ZTplbWFpbF9wcm92aWRlciBibGFja2xpc3Q6dG9rZW5zIHJlYWQ6c3RhdHMgcmVhZDppbnNpZ2h0cyByZWFkOnRlbmFudF9zZXR0aW5ncyB1cGRhdGU6dGVuYW50X3NldHRpbmdzIHJlYWQ6bG9ncyByZWFkOmxvZ3NfdXNlcnMgcmVhZDpzaGllbGRzIGNyZWF0ZTpzaGllbGRzIHVwZGF0ZTpzaGllbGRzIGRlbGV0ZTpzaGllbGRzIHJlYWQ6YW5vbWFseV9ibG9ja3MgZGVsZXRlOmFub21hbHlfYmxvY2tzIHVwZGF0ZTp0cmlnZ2VycyByZWFkOnRyaWdnZXJzIHJlYWQ6Z3JhbnRzIGRlbGV0ZTpncmFudHMgcmVhZDpndWFyZGlhbl9mYWN0b3JzIHVwZGF0ZTpndWFyZGlhbl9mYWN0b3JzIHJlYWQ6Z3VhcmRpYW5fZW5yb2xsbWVudHMgZGVsZXRlOmd1YXJkaWFuX2Vucm9sbG1lbnRzIGNyZWF0ZTpndWFyZGlhbl9lbnJvbGxtZW50X3RpY2tldHMgcmVhZDp1c2VyX2lkcF90b2tlbnMgY3JlYXRlOnBhc3N3b3Jkc19jaGVja2luZ19qb2IgZGVsZXRlOnBhc3N3b3Jkc19jaGVja2luZ19qb2IgcmVhZDpjdXN0b21fZG9tYWlucyBkZWxldGU6Y3VzdG9tX2RvbWFpbnMgY3JlYXRlOmN1c3RvbV9kb21haW5zIHVwZGF0ZTpjdXN0b21fZG9tYWlucyByZWFkOmVtYWlsX3RlbXBsYXRlcyBjcmVhdGU6ZW1haWxfdGVtcGxhdGVzIHVwZGF0ZTplbWFpbF90ZW1wbGF0ZXMgcmVhZDptZmFfcG9saWNpZXMgdXBkYXRlOm1mYV9wb2xpY2llcyByZWFkOnJvbGVzIGNyZWF0ZTpyb2xlcyBkZWxldGU6cm9sZXMgdXBkYXRlOnJvbGVzIHJlYWQ6cHJvbXB0cyB1cGRhdGU6cHJvbXB0cyByZWFkOmJyYW5kaW5nIHVwZGF0ZTpicmFuZGluZyBkZWxldGU6YnJhbmRpbmcgcmVhZDpsb2dfc3RyZWFtcyBjcmVhdGU6bG9nX3N0cmVhbXMgZGVsZXRlOmxvZ19zdHJlYW1zIHVwZGF0ZTpsb2dfc3RyZWFtcyBjcmVhdGU6c2lnbmluZ19rZXlzIHJlYWQ6c2lnbmluZ19rZXlzIHVwZGF0ZTpzaWduaW5nX2tleXMgcmVhZDpsaW1pdHMgdXBkYXRlOmxpbWl0cyBjcmVhdGU6cm9sZV9tZW1iZXJzIHJlYWQ6cm9sZV9tZW1iZXJzIGRlbGV0ZTpyb2xlX21lbWJlcnMgcmVhZDplbnRpdGxlbWVudHMgcmVhZDphdHRhY2tfcHJvdGVjdGlvbiB1cGRhdGU6YXR0YWNrX3Byb3RlY3Rpb24gcmVhZDpvcmdhbml6YXRpb25zIHVwZGF0ZTpvcmdhbml6YXRpb25zIGNyZWF0ZTpvcmdhbml6YXRpb25zIGRlbGV0ZTpvcmdhbml6YXRpb25zIGNyZWF0ZTpvcmdhbml6YXRpb25fbWVtYmVycyByZWFkOm9yZ2FuaXphdGlvbl9tZW1iZXJzIGRlbGV0ZTpvcmdhbml6YXRpb25fbWVtYmVycyBjcmVhdGU6b3JnYW5pemF0aW9uX2Nvbm5lY3Rpb25zIHJlYWQ6b3JnYW5pemF0aW9uX2Nvbm5lY3Rpb25zIHVwZGF0ZTpvcmdhbml6YXRpb25fY29ubmVjdGlvbnMgZGVsZXRlOm9yZ2FuaXphdGlvbl9jb25uZWN0aW9ucyBjcmVhdGU6b3JnYW5pemF0aW9uX21lbWJlcl9yb2xlcyByZWFkOm9yZ2FuaXphdGlvbl9tZW1iZXJfcm9sZXMgZGVsZXRlOm9yZ2FuaXphdGlvbl9tZW1iZXJfcm9sZXMgY3JlYXRlOm9yZ2FuaXphdGlvbl9pbnZpdGF0aW9ucyByZWFkOm9yZ2FuaXphdGlvbl9pbnZpdGF0aW9ucyBkZWxldGU6b3JnYW5pemF0aW9uX2ludml0YXRpb25zIHJlYWQ6b3JnYW5pemF0aW9uc19zdW1tYXJ5IGNyZWF0ZTphY3Rpb25zX2xvZ19zZXNzaW9ucyBjcmVhdGU6YXV0aGVudGljYXRpb25fbWV0aG9kcyByZWFkOmF1dGhlbnRpY2F0aW9uX21ldGhvZHMgdXBkYXRlOmF1dGhlbnRpY2F0aW9uX21ldGhvZHMgZGVsZXRlOmF1dGhlbnRpY2F0aW9uX21ldGhvZHMiLCJndHkiOiJjbGllbnQtY3JlZGVudGlhbHMifQ.NMvKpN0GUfYBGmIQvbrfeCXTE8BPud1PZPlXDlydBqYyKBOWSiXIevM8h2UZ32Yfx8CYHMAHhETIh3pkMgCSPVxojSHg9sIBsnEeZYQeOpdJ0p6eoJ-oPoFBQuWQ7dKLWy4ElDL1GQxSK_zywYf0tnIvgq5VUsQOd5BJ8Gs87vMetuxegdjq1EdPWpR7Xe_ZILVfV0M21I-QRAHlCDtcx0NFZA4vVYYjk8Q9uYGhlb47ruZWs5IUiiHM9PFEPusI2Y7ZaEMWnzWEyuPsSSKeCf7sOJcSGpTEUD0UTpi1VlSUA0S7kwm5nqpIx6LV5Nl9lFsdxZrpOzjdN8RzbcB_jg";
	auth0API, err := management.New(
		configs.EnvAuth0Domain(),
		management.WithClientCredentials(configs.EnvAuth0ClientId(), configs.EnvAuth0ClientSecret()),
		management.WithStaticToken(access_token),
	)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err}})
	}
	
	// update := &management.User{
	// 	GivenName: &user.FirstName,
	// 	FamilyName: &user.LastName,
	// 	Name: &user.Name,
	// 	UserMetadata: &map[string]interface{} {
	// 		"date_of_birth": user.DateOfBirth,
	// 		"phone": user.Phone,
	// 		"photo_file_url": user.PhotoFileUrl,
	// 		"services": user.Services,
	// 	},
	// }

	// create user object
	update := &management.User{}
	if user.FirstName != "" {
        update.GivenName = &user.FirstName;
    }
    if user.LastName != "" {
        update.FamilyName = &user.LastName;
    }
    if user.Name != "" {
        update.Name = &user.Name;
    }
	update.UserMetadata = &map[string]interface{}{}
    if user.DateOfBirth != "" {
        (*update.UserMetadata)["date_of_birth"] = &user.DateOfBirth;
    }
    if user.Phone != "" {
        (*update.UserMetadata)["phone"] = &user.DateOfBirth;
    }
    if user.PhotoFileUrl != "" {
		(*update.UserMetadata)["photo_file_url"] = user.DateOfBirth;
    }
    if user.Services != "" {
        (*update.UserMetadata)["services"] = user.DateOfBirth;
    }
	fmt.Println(update)

	err2 := auth0API.User.Update(user.UserId, update)
	fmt.Println(err2)
	if err2 != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err2}})
	}



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
	var user models.DeleteAuth0UserPayload

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	//split user ID
	user_id_by_itself := strings.Split(user.UserId,"|")[1]
	fmt.Println(user_id_by_itself)

	// check if user exists on mongo end
	var mongo_user models.Auth0User

	objId, _ := primitive.ObjectIDFromHex(user_id_by_itself)

	err := userCollection.FindOne(c.Context(), bson.M{"_id": objId}).Decode(&mongo_user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": "user does not exist"}})
	}

	//access_token := user.AccessToken;
	// get access_token for user
	access_token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IllxSnViYlM5U290UzNVdUtYVk9hbyJ9.eyJpc3MiOiJodHRwczovL2Rldi1vM25qeWhkNTRkNTJkd2Q4LnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJjOTkwYmw1T2tRcjhJNjdYNHNuMDNudzNrREtSUEoyYkBjbGllbnRzIiwiYXVkIjoiaHR0cHM6Ly9kZXYtbzNuanloZDU0ZDUyZHdkOC51cy5hdXRoMC5jb20vYXBpL3YyLyIsImlhdCI6MTY3NzYxNDQzOSwiZXhwIjoxNjc3NzAwODM5LCJhenAiOiJjOTkwYmw1T2tRcjhJNjdYNHNuMDNudzNrREtSUEoyYiIsInNjb3BlIjoicmVhZDpjbGllbnRfZ3JhbnRzIGNyZWF0ZTpjbGllbnRfZ3JhbnRzIGRlbGV0ZTpjbGllbnRfZ3JhbnRzIHVwZGF0ZTpjbGllbnRfZ3JhbnRzIHJlYWQ6dXNlcnMgdXBkYXRlOnVzZXJzIGRlbGV0ZTp1c2VycyBjcmVhdGU6dXNlcnMgcmVhZDp1c2Vyc19hcHBfbWV0YWRhdGEgdXBkYXRlOnVzZXJzX2FwcF9tZXRhZGF0YSBkZWxldGU6dXNlcnNfYXBwX21ldGFkYXRhIGNyZWF0ZTp1c2Vyc19hcHBfbWV0YWRhdGEgcmVhZDp1c2VyX2N1c3RvbV9ibG9ja3MgY3JlYXRlOnVzZXJfY3VzdG9tX2Jsb2NrcyBkZWxldGU6dXNlcl9jdXN0b21fYmxvY2tzIGNyZWF0ZTp1c2VyX3RpY2tldHMgcmVhZDpjbGllbnRzIHVwZGF0ZTpjbGllbnRzIGRlbGV0ZTpjbGllbnRzIGNyZWF0ZTpjbGllbnRzIHJlYWQ6Y2xpZW50X2tleXMgdXBkYXRlOmNsaWVudF9rZXlzIGRlbGV0ZTpjbGllbnRfa2V5cyBjcmVhdGU6Y2xpZW50X2tleXMgcmVhZDpjb25uZWN0aW9ucyB1cGRhdGU6Y29ubmVjdGlvbnMgZGVsZXRlOmNvbm5lY3Rpb25zIGNyZWF0ZTpjb25uZWN0aW9ucyByZWFkOnJlc291cmNlX3NlcnZlcnMgdXBkYXRlOnJlc291cmNlX3NlcnZlcnMgZGVsZXRlOnJlc291cmNlX3NlcnZlcnMgY3JlYXRlOnJlc291cmNlX3NlcnZlcnMgcmVhZDpkZXZpY2VfY3JlZGVudGlhbHMgdXBkYXRlOmRldmljZV9jcmVkZW50aWFscyBkZWxldGU6ZGV2aWNlX2NyZWRlbnRpYWxzIGNyZWF0ZTpkZXZpY2VfY3JlZGVudGlhbHMgcmVhZDpydWxlcyB1cGRhdGU6cnVsZXMgZGVsZXRlOnJ1bGVzIGNyZWF0ZTpydWxlcyByZWFkOnJ1bGVzX2NvbmZpZ3MgdXBkYXRlOnJ1bGVzX2NvbmZpZ3MgZGVsZXRlOnJ1bGVzX2NvbmZpZ3MgcmVhZDpob29rcyB1cGRhdGU6aG9va3MgZGVsZXRlOmhvb2tzIGNyZWF0ZTpob29rcyByZWFkOmFjdGlvbnMgdXBkYXRlOmFjdGlvbnMgZGVsZXRlOmFjdGlvbnMgY3JlYXRlOmFjdGlvbnMgcmVhZDplbWFpbF9wcm92aWRlciB1cGRhdGU6ZW1haWxfcHJvdmlkZXIgZGVsZXRlOmVtYWlsX3Byb3ZpZGVyIGNyZWF0ZTplbWFpbF9wcm92aWRlciBibGFja2xpc3Q6dG9rZW5zIHJlYWQ6c3RhdHMgcmVhZDppbnNpZ2h0cyByZWFkOnRlbmFudF9zZXR0aW5ncyB1cGRhdGU6dGVuYW50X3NldHRpbmdzIHJlYWQ6bG9ncyByZWFkOmxvZ3NfdXNlcnMgcmVhZDpzaGllbGRzIGNyZWF0ZTpzaGllbGRzIHVwZGF0ZTpzaGllbGRzIGRlbGV0ZTpzaGllbGRzIHJlYWQ6YW5vbWFseV9ibG9ja3MgZGVsZXRlOmFub21hbHlfYmxvY2tzIHVwZGF0ZTp0cmlnZ2VycyByZWFkOnRyaWdnZXJzIHJlYWQ6Z3JhbnRzIGRlbGV0ZTpncmFudHMgcmVhZDpndWFyZGlhbl9mYWN0b3JzIHVwZGF0ZTpndWFyZGlhbl9mYWN0b3JzIHJlYWQ6Z3VhcmRpYW5fZW5yb2xsbWVudHMgZGVsZXRlOmd1YXJkaWFuX2Vucm9sbG1lbnRzIGNyZWF0ZTpndWFyZGlhbl9lbnJvbGxtZW50X3RpY2tldHMgcmVhZDp1c2VyX2lkcF90b2tlbnMgY3JlYXRlOnBhc3N3b3Jkc19jaGVja2luZ19qb2IgZGVsZXRlOnBhc3N3b3Jkc19jaGVja2luZ19qb2IgcmVhZDpjdXN0b21fZG9tYWlucyBkZWxldGU6Y3VzdG9tX2RvbWFpbnMgY3JlYXRlOmN1c3RvbV9kb21haW5zIHVwZGF0ZTpjdXN0b21fZG9tYWlucyByZWFkOmVtYWlsX3RlbXBsYXRlcyBjcmVhdGU6ZW1haWxfdGVtcGxhdGVzIHVwZGF0ZTplbWFpbF90ZW1wbGF0ZXMgcmVhZDptZmFfcG9saWNpZXMgdXBkYXRlOm1mYV9wb2xpY2llcyByZWFkOnJvbGVzIGNyZWF0ZTpyb2xlcyBkZWxldGU6cm9sZXMgdXBkYXRlOnJvbGVzIHJlYWQ6cHJvbXB0cyB1cGRhdGU6cHJvbXB0cyByZWFkOmJyYW5kaW5nIHVwZGF0ZTpicmFuZGluZyBkZWxldGU6YnJhbmRpbmcgcmVhZDpsb2dfc3RyZWFtcyBjcmVhdGU6bG9nX3N0cmVhbXMgZGVsZXRlOmxvZ19zdHJlYW1zIHVwZGF0ZTpsb2dfc3RyZWFtcyBjcmVhdGU6c2lnbmluZ19rZXlzIHJlYWQ6c2lnbmluZ19rZXlzIHVwZGF0ZTpzaWduaW5nX2tleXMgcmVhZDpsaW1pdHMgdXBkYXRlOmxpbWl0cyBjcmVhdGU6cm9sZV9tZW1iZXJzIHJlYWQ6cm9sZV9tZW1iZXJzIGRlbGV0ZTpyb2xlX21lbWJlcnMgcmVhZDplbnRpdGxlbWVudHMgcmVhZDphdHRhY2tfcHJvdGVjdGlvbiB1cGRhdGU6YXR0YWNrX3Byb3RlY3Rpb24gcmVhZDpvcmdhbml6YXRpb25zIHVwZGF0ZTpvcmdhbml6YXRpb25zIGNyZWF0ZTpvcmdhbml6YXRpb25zIGRlbGV0ZTpvcmdhbml6YXRpb25zIGNyZWF0ZTpvcmdhbml6YXRpb25fbWVtYmVycyByZWFkOm9yZ2FuaXphdGlvbl9tZW1iZXJzIGRlbGV0ZTpvcmdhbml6YXRpb25fbWVtYmVycyBjcmVhdGU6b3JnYW5pemF0aW9uX2Nvbm5lY3Rpb25zIHJlYWQ6b3JnYW5pemF0aW9uX2Nvbm5lY3Rpb25zIHVwZGF0ZTpvcmdhbml6YXRpb25fY29ubmVjdGlvbnMgZGVsZXRlOm9yZ2FuaXphdGlvbl9jb25uZWN0aW9ucyBjcmVhdGU6b3JnYW5pemF0aW9uX21lbWJlcl9yb2xlcyByZWFkOm9yZ2FuaXphdGlvbl9tZW1iZXJfcm9sZXMgZGVsZXRlOm9yZ2FuaXphdGlvbl9tZW1iZXJfcm9sZXMgY3JlYXRlOm9yZ2FuaXphdGlvbl9pbnZpdGF0aW9ucyByZWFkOm9yZ2FuaXphdGlvbl9pbnZpdGF0aW9ucyBkZWxldGU6b3JnYW5pemF0aW9uX2ludml0YXRpb25zIHJlYWQ6b3JnYW5pemF0aW9uc19zdW1tYXJ5IGNyZWF0ZTphY3Rpb25zX2xvZ19zZXNzaW9ucyBjcmVhdGU6YXV0aGVudGljYXRpb25fbWV0aG9kcyByZWFkOmF1dGhlbnRpY2F0aW9uX21ldGhvZHMgdXBkYXRlOmF1dGhlbnRpY2F0aW9uX21ldGhvZHMgZGVsZXRlOmF1dGhlbnRpY2F0aW9uX21ldGhvZHMiLCJndHkiOiJjbGllbnQtY3JlZGVudGlhbHMifQ.NMvKpN0GUfYBGmIQvbrfeCXTE8BPud1PZPlXDlydBqYyKBOWSiXIevM8h2UZ32Yfx8CYHMAHhETIh3pkMgCSPVxojSHg9sIBsnEeZYQeOpdJ0p6eoJ-oPoFBQuWQ7dKLWy4ElDL1GQxSK_zywYf0tnIvgq5VUsQOd5BJ8Gs87vMetuxegdjq1EdPWpR7Xe_ZILVfV0M21I-QRAHlCDtcx0NFZA4vVYYjk8Q9uYGhlb47ruZWs5IUiiHM9PFEPusI2Y7ZaEMWnzWEyuPsSSKeCf7sOJcSGpTEUD0UTpi1VlSUA0S7kwm5nqpIx6LV5Nl9lFsdxZrpOzjdN8RzbcB_jg";
	auth0API, err := management.New(
		configs.EnvAuth0Domain(),
		management.WithClientCredentials(configs.EnvAuth0ClientId(), configs.EnvAuth0ClientSecret()),
		management.WithStaticToken(access_token),
	)
	auth0API.User.Delete(user.UserId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err}})
	}
	
	err = auth0API.User.Delete(user.UserId)
	// delete from Auth0 side
	// m.User.Delete(user.UserId)
	if err != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err}})
    }

	// delete from mongo side
	result, err := userCollection.DeleteOne(c.Context(), bson.M{"_id": objId})
	fmt.Println(result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": "user was successfully deleted"}})


}
