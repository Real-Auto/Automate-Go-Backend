package middleware

import (
	"Automate-Go-Backend/configs"
	"Automate-Go-Backend/databaseModels"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Auth0Token struct {
	AccessToken      string `json:"access_token,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	ExpiresIn        int    `json:"expires_in,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

var Now int64
var Expiration int64

var Auth0TokenVar Auth0Token

// func GetManagementApiToken(client_id string, client_secret string) (*Auth0Token, error) {
//     // ...
//     var token Auth0Token

//     // ...

//     if jsErr := json.Unmarshal(body, &token); jsErr != nil {
//         return nil, jsErr
//     }

//     return &token, nil
// }

func GetManagementApiToken() (Auth0Token, error) {
	url := configs.EnvAuth0LoginEndpoint()

	model := databaseModels.GetManagementApiTokenPayload{
		GrantType:    "client_credentials",
		ClientId:     configs.EnvAuth0ClientId(),
		ClientSecret: configs.EnvAuth0ClientSecret(),
		Audience:     configs.EnvGetManagementApiAudience(),
	}

	payload, err := json.Marshal(model)
	if err != nil {
		// Handle the error
		return Auth0Token{}, err

	}
	// fmt.Println(payload)
	// fmt.Println(bytes.NewBuffer(payload))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Add("content-type", "application/json")
	if err != nil {
		return Auth0Token{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Auth0Token{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		// handle error
		return Auth0Token{}, err
	}

	if jsErr := json.Unmarshal(body, &Auth0TokenVar); jsErr != nil {
		return Auth0Token{}, jsErr
	}

	// invalid credentials
	if Auth0TokenVar.Error != "" {
		return Auth0Token{}, errors.New("invalid credentials")
	}

	Now = time.Now().Unix()
	Expiration = Now + int64(Auth0TokenVar.ExpiresIn)

	return Auth0TokenVar, nil
}
