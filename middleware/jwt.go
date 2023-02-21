package middleware

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/form3tech-oss/jwt-go/request"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func EnsureValidToken() func(*fiber.Ctx) error {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := request.New(jwt.Keyfunc(func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		issuer := token.Claims.(jwt.MapClaims)["iss"].(string)
		issURL, err := url.Parse(issuer)
		if err != nil {
			return nil, jwt.ErrInvalidIssuer
		}
		if issURL.String() != issuerURL.String() {
			return nil, jwt.ErrInvalidIssuer
		}
		cert, err := provider.GetCertificate(token)
		if err != nil {
			return nil, err
		}
		return cert.PublicKey, nil
	}), request.WithClaims(&CustomClaims{}), request.WithIssuer(issuerURL.String()))

	errorHandler := func(c *fiber.Ctx, err error) error {
		log.Printf("Encountered error while validating JWT: %v", err)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Failed to validate JWT."})
	}

	return middleware.WithErrorHandler(func(c *fiber.Ctx, err error) error {
		if err == nil {
			// No error, continue
			return c.Next()
		}

		if _, ok := err.(*jwt.ValidationError); ok {
			// Token is invalid
			return errorHandler(c, err)
		}

		// Other errors
		return err
	})(func(c *fiber.Ctx) error {
		// Use request.ParseFromRequest() to validate the JWT in the request header
		token, err := provider.ParseFromRequest(c.Request().Request)
		if err != nil {
			return errorHandler(c, err)
		}

		// Set the user in the context
		claims := token.Claims.(*CustomClaims)
		c.Locals("user", claims)

		// Call the next handler
		return c.Next()
	})
}