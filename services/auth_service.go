package services

import (
	"example/board/auth"
	"example/board/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindCurrentUser(c *gin.Context) *models.User {
	userID, exists := c.Get(string(auth.UserIDKey))
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found UserID!"})
		return nil
	}

	var user models.User
	if err := models.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found User!"})
		return nil
	}

	return &user
}
