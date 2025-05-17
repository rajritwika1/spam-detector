package handlers

import (
	"net/http"
	"spam-detector-api/config"
	"spam-detector-api/models"

	"github.com/gin-gonic/gin"
)

// AddSpamNumber allows logged-in users to report spam numbers
func AddSpamNumber(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	authUser := user.(models.User)

	var spam models.SpamNumber
	if err := c.ShouldBindJSON(&spam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the user has already reported this number
	var existingReport models.SpamNumber
	if err := config.DB.Where("phone_number = ? AND user_id = ?", spam.PhoneNumber, authUser.ID).First(&existingReport).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "You have already reported this number"})
		return
	}

	// Save spam number with User ID
	spam.ID = authUser.ID
	if err := config.DB.Create(&spam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save spam number"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Spam number added", "data": spam})
}

// GetSpamNumbers fetches all reported spam numbers
func GetSpamNumbers(c *gin.Context) {
	var spamNumbers []models.SpamNumber
	if err := config.DB.Find(&spamNumbers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch spam numbers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": spamNumbers})
}

// CheckSpamNumber checks if a number is reported as spam
func CheckSpamNumber(c *gin.Context) {
	phoneNumber := c.Param("phoneNumber")

	var spam models.SpamNumber
	if err := config.DB.Where("phone_number = ?", phoneNumber).First(&spam).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Number not found in spam list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Spam number found", "data": spam})
}
