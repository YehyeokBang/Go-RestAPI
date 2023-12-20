package controllers

import (
	"net/http"

	"example/board/models"
	"example/board/types"

	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {
	user := FindCurrentUser(c)

	responseUser := types.ResponseUser{
		ID:       user.ID,
		Name:     user.Name,
		Nickname: user.Nickname,
		Email:    user.Email,
	}
	c.JSON(http.StatusOK, gin.H{"data": responseUser})
}

func CreateUser(c *gin.Context) {
	var input types.RequestCreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := models.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password!"})
		return
	}

	user := models.User{
		Name:     input.Name,
		Nickname: input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}
	models.DB.Create(&user)

	responseUser := types.ResponseUser{
		ID:       user.ID,
		Name:     user.Name,
		Nickname: user.Nickname,
		Email:    user.Email,
	}
	c.JSON(http.StatusOK, gin.H{"data": responseUser})
}

func UpdateUser(c *gin.Context) {
	currentUser := FindCurrentUser(c)

	var input types.RequestUpdateUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		models.DB.Model(&currentUser).Update("Name", input.Name)
	}
	if input.Nickname != "" {
		models.DB.Model(&currentUser).Update("Nickname", input.Nickname)
	}
	if input.Email != "" {
		models.DB.Model(&currentUser).Update("Email", input.Email)
	}

	responseUser := types.ResponseUser{
		ID:       currentUser.ID,
		Name:     currentUser.Name,
		Nickname: currentUser.Nickname,
		Email:    currentUser.Email,
	}

	c.JSON(http.StatusOK, gin.H{"data": responseUser})
}

func DeleteUser(c *gin.Context) {
	currentUser := FindCurrentUser(c)

	models.DB.Where("user_id = ?", currentUser.ID).Delete(&models.Post{})
	models.DB.Delete(&currentUser)

	c.JSON(http.StatusOK, gin.H{"data": "User deleted!"})
}
