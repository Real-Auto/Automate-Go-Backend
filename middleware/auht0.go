package middleware

import (
	"Automate-Go-Backend/configs"
	"Automate-Go-Backend/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

type Auth0Token struct {
    AccessToken string `json:"access_token,omitempty"`
    TokenType   string `json:"token_type,omitempty"`
    ExpiresIn   int    `json:"expires_in,omitempty"`
	Error 		string `json:"error,omitempty"`
	ErrorDescription		string `json:"error_description,omitempty"`
}

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

	model := models.GetManagementApiTokenPayload {
		GrantType: "client_credentials",
		ClientId: "sdf",
		ClientSecret: configs.EnvAuth0ClientSecret(),
		Audience: configs.EnvGetManagementApiAudience(),
	}

	payload, err := json.Marshal(model)
	if err != nil {
		// Handle the error
		return Auth0Token{}, err
		
	}
	fmt.Println(payload)
	fmt.Println(bytes.NewBuffer(payload))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
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
	if err != nil {
		// handle error
		return Auth0Token{}, err
	}	

	var responseData Auth0Token;
	if jsErr := json.Unmarshal(body, &responseData); jsErr != nil {
		return Auth0Token{}, jsErr
	}

	// // invalid credentials
	// if (responseData.Error != "") {
	// 	return Auth0Token{}, json.UnmarshalTypeError{responseData}
	// }
	// fmt.Println(res)
	fmt.Println(string(body))

	return responseData, nil
}