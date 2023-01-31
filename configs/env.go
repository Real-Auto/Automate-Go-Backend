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

func EnvAuth0ClientID() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("AUTH0_CLIENT_ID")
}

func EnvAuth0Connection() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("CONNECTION")
}

func EnvAuth0SignupEndpoint() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("AUTH0_SIGN_UP_ENDPOINT")
}



