package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Write functions for getting environment variables here

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}

func EnvAuth0ClientId() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("AUTH0_CLIENT_ID")
}

func EnvAuth0ClientSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("AUTH0_CLIENT_SECRET")
}

func EnvAuth0Connection() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("CONNECTION")
}

func EnvAuth0ApiAudience() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("API_AUDIENCE")
}

func EnvAuth0GetUserInfoEndpoint() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("AUTH0_GET_USER_INFO_ENDPOINT")
}

func EnvAuth0SignupEndpoint() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("AUTH0_SIGN_UP_ENDPOINT")
}

func EnvAuth0LoginEndpoint() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("AUTH0_LOGIN_ENDPOINT")
}

func EnvAuth0ChangePasswordEndpoint() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("AUTH0_CHANGE_PASSWORD_ENDPOINT")
}

func EnvAuth0UpdateUserEndpoint() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("AUTH0_UPDATE_USER_ENDPOINT")
}

func EnvAuth0DeleteUserEndpoint() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("AUTH0_DELETE_USER_ENDPOINT")
}
