package controllers

import (
	"net/http"

	"example/board/services"
	"example/board/types"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func (controller *UserController) FindUser(c *gin.Context) {
	responseUser, err := controller.UserService.FindUserByID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responseUser})
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var input types.RequestCreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseUser, err := controller.UserService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responseUser})
}

func (controller *UserController) UpdateUser(c *gin.Context) {
	var input types.RequestUpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseUser, err := controller.UserService.UpdateUser(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responseUser})
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	err := controller.UserService.DeleteUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User deleted!"})
}
