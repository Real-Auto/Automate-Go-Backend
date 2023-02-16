package controllers

import (
	"Automate-Go-Backend/configs"
	"Automate-Go-Backend/models"
	"Automate-Go-Backend/responses"
	"context"
	// "fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

// func CreateUser(c *fiber.Ctx) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	var user models.User
// 	defer cancel()

// 	//validate the request body
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	//use the validator library to validate required fields
// 	if validationErr := validate.Struct(&user); validationErr != nil {
// 		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
// 	}

// 	newUser := models.User{
// 		FirstName:    user.FirstName,
// 		LastName:     user.LastName,
// 		DateOfBirth:  user.DateOfBirth,
// 		Phone:        user.Phone,
// 		PhotoFileUrl: user.PhotoFileUrl,
// 		Services:     user.Services,
// 		Email:        user.Email,
// 		Password:     user.Password,
// 	}

// 	result, err := userCollection.InsertOne(ctx, newUser)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
// }


// func EditProfileInformation(c *fiber.Ctx) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	userId := c.Params("userId")
// 	var user models.EditProfileInformationPayload
// 	defer cancel()

// 	objId, _ := primitive.ObjectIDFromHex(userId)
// 	fmt.Println(objId)
// 	//validate the request body
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
// 	}
// 	fmt.Println(user)

// 	//use the validator library to validate required fields
// 	if validationErr := validate.Struct(&user); validationErr != nil {
// 		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
// 	}

// 	// retrieve current meta_data information
// 	var currentUser models.Auth0User
// 	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&currentUser)

// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
// 	}
	
// 	update := bson.M{}
//     user_meta_data := bson.M{}
//     update["user_metadata"] = user_meta_data

// 	if user.FirstName != "" {
// 		fmt.Println("hello chap1")
// 		update["given_name"] = user.FirstName
// 	}
// 	if user.FirstName != "" {
// 		update["family_name"] = user.LastName
// 	} 
// 	if user.Name != "" {
// 		update["name"] = user.Name
// 	} 
// 	if user.DateOfBirth != "" {
// 		currentUser.MetaData.DateOfBirth = user.DateOfBirth
// 	} 
// 	if user.Phone != "" {
// 		currentUser.MetaData.Phone = user.Phone
// 	} 
// 	if user.PhotoFileUrl != "" {
// 		currentUser.MetaData.PhotoFileUrl = user.PhotoFileUrl
// 	} 
// 	if user.Services != "" {
// 		currentUser.MetaData.Services = user.Services
// 	}
// 	update["user_metadata"] = currentUser.MetaData

	
	
	

		
// 	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
// 	fmt.Println(update)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	//get updated user details
// 	var updatedUser models.Auth0User
	
// 	if result.MatchedCount == 1 {
// 		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)

// 		if err != nil {
// 			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
// 		}
// 	}

// 	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedUser}})
// }

func DeleteAUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
	)
}

func GetAllUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.Auth0User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.Auth0User
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		users = append(users, singleUser)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": users}},
	)
}
