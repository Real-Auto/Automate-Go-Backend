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
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   int    `json:"expires_in"`
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

func GetManagementApiToken(client_id string, client_secret string) (map[string]interface{}, error) {
	url := configs.EnvAuth0LoginEndpoint()

	model := models.GetManagementApiTokenPayload {
		GrantType: "client_credentials",
		ClientId: configs.EnvAuth0ClientId(),
		ClientSecret: configs.EnvAuth0ClientSecret(),
		Audience: configs.EnvGetManagementApiAudience(),
	}

	payload, err := json.Marshal(model)
	if err != nil {
		// Handle the error
		return make(map[string]interface{}), err
		
	}
	fmt.Println(payload)
	fmt.Println(bytes.NewBuffer(payload))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Add("content-type", "application/json")
	if err != nil {
		return make(map[string]interface{}), err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return make(map[string]interface{}), err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var responseData map[string]interface{}
	if jsErr := json.Unmarshal(body, &responseData); jsErr != nil {
		return make(map[string]interface{}), jsErr
	}
	if err != nil {
		// handle error
		return make(map[string]interface{}), err
	}	
	
	// fmt.Println(res)
	fmt.Println(string(body))

	return responseData, nil
}