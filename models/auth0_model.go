package models

import (
	"time"

)

// swagger:model
type UserMetaData struct {
	Services     string 	`bson:"services" json:"services" validate:"required"`
	DateOfBirth  string 	`bson:"date_of_birth" json:"date_of_birth" validate:"required"`
	PhotoFileUrl string 	`bson:"photo_file_url" json:"photo_file_url" validate:"required"`
	Phone  		 string 	`bson:"phone" json:"phone" validate:"required"`
	Language	 string 	`bson:"language" json:"language" validate:"required"`
	CreationDate time.Time 	`bson:"creation_date" json:"creation_date" validate:"required"`
	LastUpdated	 time.Time 	`bson:"last_updated" json:"last_updated" validate:"required"`
}

// swagger:model
type Auth0User struct {
	ID           string       `bson:"_id,omitempty" json:"_id,omitempty"`
	Tenant       string       `bson:"tenant,omitempty" json:"tenant,omitempty"`
	ClientId     string       `bson:"client_id" json:"client_id" validate:"required"`
	Connection   string       `bson:"connection" json:"connection" validate:"required"`
	Email        string       `bson:"email" json:"email" validate:"required"`
	Password     string       `bson:"password" json:"password" validate:"required"`
	Name         string       `bson:"name" json:"name" validate:"required"`
	GivenName    string       `bson:"given_name" json:"given_name" validate:"required"`
	FamilyName   string       `bson:"family_name" json:"family_name" validate:"required"`
	UserMetaData UserMetaData `bson:"user_metadata" json:"user_metadata" validate:"required"`
}

// swagger:model
type Auth0UserLogin struct {
	GrantType    string `json:"grant_type" validate:"required"`
	ClientId     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
	Audience     string `json:"audience" validate:"required"`
	Email        string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Scope        string `json:"scope,omitempty"`
	Realm        string `json:"realm,omitempty"`
}

// swagger:model
type Auth0UserChangePassword struct {
	ClientId   string `json:"client_id" validate:"required"`
	Email      string `json:"username" validate:"required"`
	Connection string `json:"connection" validate:"required"`
}

// swagger:model
type SignUpPayload struct {
	FirstName    string 	`bson:"first_name" json:"first_name" validate:"required"`
	LastName     string 	`bson:"last_name" json:"last_name" validate:"required"`
	Name         string 	`bson:"name" json:"name" validate:"required"`
	DateOfBirth  string 	`bson:"date_of_birth" json:"date_of_birth" validate:"required"`
	Phone        string 	`bson:"phone" json:"phone" validate:"required"`
	PhotoFileUrl string 	`bson:"photo_file_url" json:"photo_file_url" validate:"required"`
	Services     string 	`bson:"services" json:"services" validate:"required"`
	Language	 string 	`bson:"language" json:"language" validate:"required"`
	Email        string 	`bson:"email" json:"email" validate:"required"`
	Password     string 	`bson:"password" json:"password" validate:"required"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// swagger:model
type ChangePasswordPayload struct {
	Email string `json:"email" validate:"required"`
}

// swagger:model
type GetAuth0UserFieldsPayload struct {
    AccessToken string `json:"access_token" validate:"required"`
}

type GetAuth0UserResponse struct {
	Sub 			string		`bson:"sub" json:"sub" validate:"required"`
    GivenName 		string 		`bson:"given_name" json:"given_name" validate:"required"`
    FamilyName 		string		`bson:"family_name" json:"family_name" validate:"required"`	
    Nickname		string		`bson:"nickname" json:"nickname" validate:"required"`
    Name 			string		`bson:"name" json:"name" validate:"required"`
    Picture			string		`bson:"picture" json:"picture" validate:"required"`
    UpdatedAt 		time.Time	`bson:"updated_at" json:"updated_at" validate:"required"`
    Email 			string		`bson:"email" json:"email" validate:"required"`
	EmailVerified	bool		`bson:"email_verified" json:"email_verified" validate:"required"`
    Services 		string		`bson:"services" json:"services" validate:"required"`
    DateOfBirth 	string		`bson:"date_of_birth" json:"date_of_birth" validate:"required"`
    PhotoFileUrl	string		`bson:"photo_file_url" json:"photo_file_url" validate:"required"`
    Phone			string		`bson:"phone" json:"phone" validate:"required"`
	Language	 	string		`bson:"language" json:"language" validate:"required"`
	CreationDate	time.Time 	`bson:"creation_date" json:"creation_date" validate:"required"`
	LastUpdated	 	time.Time	`bson:"last_updated" json:"last_updated" validate:"required"`
} 

type DeleteAuth0UserPayload struct {
	UserId string `bson:"user_id" json:"user_id" validate:"required"`
	AccessToken string `bson:"access_token" json:"access_token"`
}

type UpdateAuth0UserPayload struct {
	UserId string `bson:"user_id" json:"user_id" validate:"required"`
	FirstName    string `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName     string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Name         string `bson:"name,omitempty" json:"name,omitempty"`
	DateOfBirth  string `bson:"date_of_birth,omitempty" json:"date_of_birth,omitempty"`
	Phone        string `bson:"phone,omitempty" json:"phone,omitempty"`
	PhotoFileUrl string `bson:"photo_file_url,omitempty" json:"photo_file_url,omitempty"`
	Services     string `bson:"services,omitempty" json:"services,omitempty"`
	Language	 string `bson:"language,omitempty" json:"language,omitempty" validate:"required"`
}

type GetManagementApiTokenPayload struct {
	GrantType string `bson:"grant_type" json:"grant_type" validate:"required"`
	ClientId string `bson:"client_id" json:"client_id" validate:"required"`
	ClientSecret string `bson:"client_secret" json:"client_secret" validate:"required"`
	Audience string `bson:"audience" json:"audience" validate:"required"`
}

