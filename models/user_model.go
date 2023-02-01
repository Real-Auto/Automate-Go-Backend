package models

//import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	FirstName   string `bson:"first_name" validate:"required"`
	LastName  string `bson:"last_name" validate:"required"`
	DateOfBirth   string `bson:"date_of_birth" validate:"required"`
	Phone  string `bson:"phone" validate:"required"`
	PhotoFileUrl  string `bson:"photo_file_url" validate:"required"`
	Services  string `bson:"services" validate:"required"`
	Email  string `bson:"email" validate:"required"`
	Password string `bson:"password" validate:"required"`
}

