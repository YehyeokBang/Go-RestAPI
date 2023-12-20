package controllers

import (
	"net/http"

	"example/board/models"
	"example/board/types"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	user := FindCurrentUser(c)

	var input types.RequestCreatePost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content, UserID: user.ID}
	models.DB.Create(&post)

	models.DB.Preload("User").First(&post)

	responsePost := types.ResponsePost{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		UserID:   post.UserID,
		Nickname: post.User.Nickname,
	}
	c.JSON(http.StatusOK, gin.H{"data": responsePost})
}

func FindPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Preload("User").Find(&posts)

	var responsePosts []types.ResponsePost
	for _, post := range posts {
		responsePost := types.ResponsePost{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			UserID:   post.UserID,
			Nickname: post.User.Nickname,
		}
		responsePosts = append(responsePosts, responsePost)
	}
	c.JSON(http.StatusOK, gin.H{"data": responsePosts})
}

func FindPost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Preload("User").Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found Post!"})
		return
	}

	responsePost := types.ResponsePost{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		UserID:   post.UserID,
		Nickname: post.User.Nickname,
	}
	c.JSON(http.StatusOK, gin.H{"data": responsePost})
}

func UpdatePost(c *gin.Context) {
	var post models.Post

	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found Post!"})
		return
	}

	currentUser := FindCurrentUser(c)
	if post.UserID != currentUser.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Match User!"})
		return
	}

	var input types.RequestUpdatePost

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title != "" {
		models.DB.Model(&post).Update("Title", input.Title)
	}
	if input.Content != "" {
		models.DB.Model(&post).Update("Content", input.Content)
	}

	responsePost := types.ResponsePost{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		UserID:   post.UserID,
		Nickname: post.User.Nickname,
	}
	c.JSON(http.StatusOK, gin.H{"data": responsePost})
}

func DeletePost(c *gin.Context) {
	var post models.Post

	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found Post!"})
		return
	}

	currentUser := FindCurrentUser(c)
	if post.UserID != currentUser.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Match User!"})
		return
	}

	models.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"data": "Post deleted!"})
}
