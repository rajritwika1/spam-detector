package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadSecretKey loads the SECRET_KEY from .env
func LoadSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY is not set in the environment")
	}
	return secretKey
}
