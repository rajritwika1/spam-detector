package utils

import (
	"spam-detector-api/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte(config.LoadSecretKey())

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
