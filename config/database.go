package config

import (
	"fmt"
	"log"
	"os"
	"spam-detector-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading.env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Contact{}, &models.SpamNumber{})
	if err != nil {
		log.Fatal("Failed to auto-migrate models:", err)
	}

	DB = db
	fmt.Println("Databse connected and migrated Successfully!")
}
