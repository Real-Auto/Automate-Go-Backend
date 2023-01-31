package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auth0User struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Tenant      string             `bson:"tenant"`
    ClientID    string             `bson:"client_id" validate:"required"`
    Connection  string             `bson:"connection" validate:"required"`
    Email       string             `bson:"email" validate:"required,email"`
    Password    string             `bson:"password" validate:"required"`
	GivenName   string `bson:"given_name" validate:"required"`
	FamilyName  string `bson:"family_name" validate:"required"`
	EmailVerified bool  `bson:"email_verified" validate:"required"`
    MetaData  map[string]string `bson:"metadata" validate:"required"`
}