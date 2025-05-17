package routes

import (
	"spam-detector-api/controllers"
	"spam-detector-api/handlers"
	"spam-detector-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/auth/register", controllers.Register)
		api.POST("/auth/login", controllers.Login)

		spam := api.Group("/spam", middleware.AuthMiddleware()) // Protected routes
		{
			spam.POST("/", handlers.AddSpamNumber)              // Add a spam number
			spam.GET("/", handlers.GetSpamNumbers)              // Get all spam numbers
			spam.GET("/:phoneNumber", handlers.CheckSpamNumber) // Check if a number is spam
		}

		contact := api.Group("/contacts", middleware.AuthMiddleware()) // Protected routes
		{
			contact.POST("/", handlers.AddContact)         // Add a contact
			contact.GET("/", handlers.GetContacts)         // Get all contacts
			contact.DELETE("/:id", handlers.DeleteContact) // Delete a contact
		}
	}
}
