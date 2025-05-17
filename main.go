package main

import (
	"fmt"
	"log"
	"os"
	"spam-detector-api/config"
	"spam-detector-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Check if secret key is available
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY is missing in .env")
	}
	config.InitDB()

	router := gin.Default()
	routes.SetupRoutes(router)

	port := ":8080"

	fmt.Println("Server running on port 8080")
	//log.Fatal(r.Run(":8080"))
	router.Run(port)
}
