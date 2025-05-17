package handlers

import (
	"net/http"
	"spam-detector-api/config"
	"spam-detector-api/models"

	"github.com/gin-gonic/gin"
)

func DummyHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Dummy handler for Swagger!"})
}

// AddContact allows a user to add a new contact
func AddContact(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	authUser := user.(models.User)

	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if contact already exists for the user
	var existingContact models.Contact
	if err := config.DB.Where("phone_number = ? AND user_id = ?", contact.PhoneNumber, authUser.ID).First(&existingContact).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Contact already exists"})
		return
	}

	// Assign the user ID to the contact and save
	contact.UserID = authUser.ID
	if err := config.DB.Create(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save contact"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Contact added", "data": contact})
}

// GetContacts fetches all contacts of the logged-in user
func GetContacts(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	authUser := user.(models.User)
	var contacts []models.Contact
	if err := config.DB.Where("user_id = ?", authUser.ID).Find(&contacts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contacts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}

// DeleteContact removes a contact by ID
func DeleteContact(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	authUser := user.(models.User)
	contactID := c.Param("id")

	// Find the contact and ensure it belongs to the user
	var contact models.Contact
	if err := config.DB.Where("id = ? AND user_id = ?", contactID, authUser.ID).First(&contact).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	// Delete the contact
	if err := config.DB.Delete(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}
