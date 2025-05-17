package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	PhoneNumber string `gorm:"unique;not null"`
	Email       string `gorm:"unique"`
	UserID      uint   `gorm:"not null"` 
	User        User   `gorm:"foreignKey:UserID"`
}
