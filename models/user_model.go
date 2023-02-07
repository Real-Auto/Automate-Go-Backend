package models

//import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	FirstName    string `bson:"first_name" json:"first_name" validate:"required"`
	LastName     string `bson:"last_name" json:"last_name" validate:"required"`
	DateOfBirth  string `bson:"date_of_birth" json:"date_of_birth" validate:"required"`
	Phone        string `bson:"phone" json:"phone" validate:"required"`
	PhotoFileUrl string `bson:"photo_file_url" json:"photo_file_url" validate:"required"`
	Services     string `bson:"services" json:"services" validate:"required"`
	Email        string `bson:"email" json:"email" validate:"required"`
	Password     string `bson:"password" json:"password" validate:"required"`
}

type EditProfileInformationPayload struct {
	FirstName    string `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName     string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	DateOfBirth  string `bson:"date_of_birth,omitempty" json:"date_of_birth,omitempty"`
	Phone        string `bson:"phone,omitempty" json:"phone,omitempty"`
	PhotoFileUrl string `bson:"photo_file_url,omitempty" json:"photo_file_url,omitempty" `
	Services     string `bson:"services,omitempty" json:"services,omitempty"`
}
