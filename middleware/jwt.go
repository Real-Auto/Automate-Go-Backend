package middleware

import (
	"Automate-Go-Backend/configs"
	"Automate-Go-Backend/responses"

	// "context"
	"strings"

	// "encoding/json"
	//"time"
	"encoding/json"
	//"errors"
	"io/ioutil"

	// "reflect"
	//"bytes"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	// "github.com/auth0-community/auth0"
	// "github.com/auth0-community/auth0/auth"
	// "github.com/auth0-community/go-auth0"

	// "github.com/auth0/go-auth0"
	// "github.com/auth0/go-auth0/management"

	// "encoding/json"
	// "errors"
	"fmt"
	// "io/ioutil"
	"net/http"
)

// EnsureValidToken is a middleware that will check the validity of our JWT.
func ValidateToken(funcScopes []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
	
		if tokenString == "" {
			// The Authorization header is not provided
			return c.Status(fiber.StatusUnauthorized).JSON(responses.UserResponse{
				Status: http.StatusUnauthorized,
				Message: "error",
				Data:   &fiber.Map{"data": "Authorization Header is Missing"},
			})
		}

		tokenString = strings.Split(tokenString, " ")[1]


		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Retrieve the public key from Auth0
			cert, err := getPemCert(c, configs.EnvAuth0DomainByItself(), configs.EnvAuth0ApiAudience(), token.Header["kid"].(string))
			if err != nil {
				return nil, err
			}

			// Parse the public key
			//fmt.Println(cert)
			parsedCert, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			if err != nil {
				return nil, err
			}

			// Return the parsed public key
			return parsedCert, nil
		})

		if err != nil {
			fmt.Println(err)
			fmt.Println("heloogosegjse")
			return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
				Status: http.StatusUnauthorized,
				Data:   &fiber.Map{"data": err},
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
				Status: http.StatusUnauthorized,
				Data:   &fiber.Map{"data": err},
			})
		}

		// Validate the audience
		aud := configs.EnvAuth0ApiAudience()
		if !claims.VerifyAudience(aud, false) {
			return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
				Status: http.StatusUnauthorized,
				Data:   &fiber.Map{"data": "Invalid Audience"},
			})
		}

		// Validate the scope
		scopes, ok := claims["scope"].(string)

		scopes_arr := strings.Fields(scopes)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
				Status: http.StatusUnauthorized,
				Data:   &fiber.Map{"data": "Invalid Scope"},
			})
		}

		requiredScopes := funcScopes
		for _, scope := range requiredScopes {
			found := false
			for _, s := range scopes_arr {
				if s == scope {
					found = true
					break
				}
			}

			if !found {
				return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
					Status: http.StatusUnauthorized,
					Message: "error",
					Data:   &fiber.Map{"data": "Insufficient Scope"},
				})
			}
		}

		// Store the user ID in the context
		c.Locals("userID", claims["sub"])

		// Continue with the next handler
		return c.Next()
	}
}

func getPemCert(c *fiber.Ctx, domain string, audience string, kid string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/.well-known/jwks.json", domain), nil)
	// fmt.Println(fmt.Sprintf("https://%s/.well-known/jwks.json", domain))
	if err != nil {
		return "", c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
			Status: http.StatusUnauthorized,
			Message: "error",
			Data:   &fiber.Map{"data": err},
		})
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Charset", "utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return "", c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
			Status: http.StatusUnauthorized,
			Message: "error",
			Data:   &fiber.Map{"data": err},
		})
	}

	defer resp.Body.Close()

	// Parse the JWKS into a map
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
			Status: http.StatusUnauthorized,
			Message: "error",
			Data:   &fiber.Map{"data": err},
		})
	}
	var jwks map[string]interface{}
	//fmt.Println(bytes.NewBuffer(body))
	if err := json.Unmarshal(body, &jwks); err != nil {
		return "", c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
			Status: http.StatusUnauthorized,
			Message: "error",
			Data:   &fiber.Map{"data": err},
		})
	}

	// Find the public key with the specified key ID (kid)
	var cert string;
	keys, ok := jwks["keys"].([]interface{})
	if !ok {
		return "", c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
			Status: http.StatusUnauthorized,
			Message: "error",
			Data:   &fiber.Map{"data": "invalid JWKS format"},
		})
	}
	for _, key := range keys {
		k, ok := key.(map[string]interface{})
		if !ok {
			return "", c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
				Status: http.StatusUnauthorized,
				Message: "error",
				Data:   &fiber.Map{"data": "invalid JWKS format"},
			})
		}
		if k["kid"].(string) == kid {
			alg := k["alg"].(string)
			if alg != "RS256" {
				return "", c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
					Status: http.StatusUnauthorized,
					Message: "error",
					Data:   &fiber.Map{"data": fmt.Errorf("unsupported algorithm: %s", alg)},
				})
			}
			cert = k["x5c"].([]interface{})[0].(string)
			break
		}
	}
	if cert == "" {
		return "", c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
			Status: http.StatusUnauthorized,
			Message: "error",
			Data:   &fiber.Map{"data": fmt.Errorf("public key not found for kid: %s", kid)},
		})
	}

	// Format the public key certificate as a PEM block
	pemCert := "-----BEGIN CERTIFICATE-----\n" + cert + "\n-----END CERTIFICATE-----\n"

	return pemCert, nil
}
