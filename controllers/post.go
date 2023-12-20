package controllers

import (
	"net/http"
	"strconv"

	"example/board/models"
	"example/board/services"
	"example/board/types"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	PostService services.PostService
}

func (controller *PostController) CreatePost(c *gin.Context) {
	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responsePost, err := controller.PostService.CreatePost(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responsePost})
}

func (controller *PostController) FindPosts(c *gin.Context) {
	responsePosts, err := controller.PostService.FindPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find posts!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responsePosts})
}

func (controller *PostController) FindPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid postID format!"})
		return
	}

	responsePost, err := controller.PostService.FindPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find post!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responsePost})
}

func (controller *PostController) UpdatePost(c *gin.Context) {
	var input types.RequestUpdatePost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responsePost, err := controller.PostService.UpdatePost(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responsePost})
}

func (controller *PostController) DeletePost(c *gin.Context) {
	err := controller.PostService.DeletePost(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Post deleted!"})
}
