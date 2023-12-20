package services

import (
	"example/board/models"
	"example/board/models/repositories"
	"example/board/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostService struct {
	PostRepository *repositories.PostRepository
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{
		PostRepository: repositories.NewPostRepository(db),
	}
}

func (service *PostService) CreatePost(c *gin.Context, input models.Post) (types.ResponsePost, error) {
	currentUser := FindCurrentUser(c)

	inputToCreate := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  currentUser.ID,
	}

	newPost, err := service.PostRepository.Create(&inputToCreate)
	if err != nil {
		return types.ResponsePost{}, err
	}

	responsePost := types.ResponsePost{
		ID:       newPost.ID,
		Title:    newPost.Title,
		Content:  newPost.Content,
		UserID:   newPost.UserID,
		Nickname: newPost.User.Nickname,
	}
	return responsePost, nil
}

func (service *PostService) FindPosts() ([]types.ResponsePost, error) {
	posts, err := service.PostRepository.FindAll()
	if err != nil {
		return nil, err
	}

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
	return responsePosts, nil
}

func (service *PostService) FindPostByID(id uint) (types.ResponsePost, error) {
	post, err := service.PostRepository.FindByID(id)
	if err != nil {
		return types.ResponsePost{}, err
	}

	responsePost := types.ResponsePost{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		UserID:   post.UserID,
		Nickname: post.User.Nickname,
	}
	return responsePost, nil
}

func (service *PostService) UpdatePost(c *gin.Context, input types.RequestUpdatePost) (types.ResponsePost, error) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		return types.ResponsePost{}, err
	}

	currentUser := FindCurrentUser(c)
	if currentUser.ID != post.UserID {
		return types.ResponsePost{}, nil
	}

	if input.Title != "" {
		post.Title = input.Title
	}
	if input.Content != "" {
		post.Content = input.Content
	}

	updatedPost, err := service.PostRepository.Update(&post)
	if err != nil {
		return types.ResponsePost{}, err
	}

	responsePost := types.ResponsePost{
		ID:       updatedPost.ID,
		Title:    updatedPost.Title,
		Content:  updatedPost.Content,
		UserID:   updatedPost.UserID,
		Nickname: updatedPost.User.Nickname,
	}
	return responsePost, nil
}

func (service *PostService) DeletePost(c *gin.Context) error {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		return err
	}

	currentUser := FindCurrentUser(c)
	if currentUser.ID != post.UserID {
		return nil
	}

	if err := service.PostRepository.Delete(&post); err != nil {
		return err
	}
	return nil
}
