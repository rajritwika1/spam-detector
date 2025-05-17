package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name"`
	PhoneNumber string `gorm:"unique" json:"phone_number"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"-"`
}
type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}
