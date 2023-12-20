package controllers

import (
	"fmt"
	"net/http"
	"os"

	"example/board/auth"
	"example/board/models"
	"example/board/types"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Login(c *gin.Context) {
	var input types.RequestLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email!"})
		return
	}

	if err := models.VerifyPassword(user.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password!"})
		return
	}

	loadErr := godotenv.Load(".env")
	if loadErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load env file!"})
	}

	secret := os.Getenv("JWT_SECRET")

	claim := auth.NewClaim(fmt.Sprint(user.ID))

	token, err := auth.GenerateToken(claim, secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token!"})
		return
	}

	c.JSON(http.StatusOK, types.ResponseToken{Token: token})
}

func FindCurrentUser(c *gin.Context) models.User {
	userID, exists := c.Get(string(auth.UserIDKey))
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found UserID!"})
		return models.User{}
	}

	var user models.User
	if err := models.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found User!"})
		return models.User{}
	}

	return user
}
