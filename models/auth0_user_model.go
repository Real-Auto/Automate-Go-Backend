package models

type UserMetaData struct {
	Services      string `bson:"services" json:"services" validate:"required"`
	DateOfBirth  string `bson:"date_of_birth" json:"date_of_birth" validate:"required"`
	PhotoFileUrl     string `bson:"photo_file_url" json:"photo_file_url" validate:"required"`
	Phone  string `bson:"phone" json:"phone" validate:"required"`
}


type Auth0User struct {
    ID          string `bson:"_id,omitempty" json:"_id,omitempty"`
    Tenant      string             `bson:"tenant,omitempty" json:"tenant,omitempty"`
    ClientId    string             `bson:"client_id" json:"client_id" validate:"required"`
    Connection  string             `bson:"connection" json:"connection" validate:"required"`
    Email       string             `bson:"email" json:"email" validate:"required"`
    Password    string             `bson:"password" json:"password" validate:"required"`
	GivenName   string `bson:"given_name" json:"given_name" validate:"required"`
	FamilyName  string `bson:"family_name" json:"family_name" validate:"required"`
    MetaData   UserMetaData `bson:"user_metadata" json:"user_metadata" validate:"required"`
}



type Auth0UserLogin struct {
    GrantType string `json:"grant_type" validate:"required"`
    ClientId string `json:"client_id" validate:"required"`
    ClientSecret string `json:"client_secret" validate:"required"`
    Audience string `json:"audience" validate:"required"`
    Email string `json:"username" validate:"required"`
    Password string `json:"password" validate:"required"`
    Scope string `json:"scope,omitempty"`
    Realm string `json:"realm,omitempty"`

}

type Auth0UserChangePassword struct {
    ClientId    string  `json:"client_id" validate:"required"`
    Email string `json:"username" validate:"required"`
    Connection  string  `json:"connection" validate:"required"`

}

type LoginPayload struct {
	Email string `json:"email" validate:"required"`
    Password string `json:"password" validate:"required"`
}

type ChangePasswordPayload struct {
    Email string `json:"email" validate:"required"`
}