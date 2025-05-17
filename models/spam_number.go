package models

import (
	"time"

	"gorm.io/gorm"
)

type SpamNumber struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	PhoneNumber string    `gorm:"unique;not null"`
	Reason      string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
