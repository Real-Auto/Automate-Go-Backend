package models

type Auth0User struct {
    ID          string `json:"_id,omitempty"`
    Tenant      string             `json:"tenant,omitempty"`
    ClientId    string             `json:"client_id" validate:"required"`
    Connection  string             `json:"connection" validate:"required"`
    Email       string             `json:"email" validate:"required,email"`
    Password    string             `json:"password" validate:"required"`
	GivenName   string `json:"given_name" validate:"required"`
	FamilyName  string `json:"family_name" validate:"required"`
    MetaData  map[string]string `json:"meta_data" validate:"required"`
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