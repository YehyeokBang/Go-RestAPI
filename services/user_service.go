package services

import (
	"example/board/models"
	"example/board/models/repositories"
	"example/board/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository *repositories.UserRepository
	PostRepository *repositories.PostRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		UserRepository: repositories.NewUserRepository(db),
	}
}

func (service *UserService) CreateUser(input types.RequestCreateUser) (types.ResponseUser, error) {
	hashedPassword, err := models.HashPassword(input.Password)
	if err != nil {
		return types.ResponseUser{}, err
	}

	userToCreate := models.User{
		Name:     input.Name,
		Nickname: input.Name, // 기본값으로 이름을 닉네임으로 설정
		Email:    input.Email,
		Password: hashedPassword,
	}

	newUser, err := service.UserRepository.Create(&userToCreate)
	if err != nil {
		return types.ResponseUser{}, err
	}

	responseUser := types.ResponseUser{
		ID:       newUser.ID,
		Name:     newUser.Name,
		Nickname: newUser.Nickname,
		Email:    newUser.Email,
	}
	return responseUser, nil
}

func (service *UserService) FindUserByID(c *gin.Context) (types.ResponseUser, error) {
	user := FindCurrentUser(c)

	responseUser := types.ResponseUser{
		ID:       user.ID,
		Name:     user.Name,
		Nickname: user.Nickname,
		Email:    user.Email,
	}
	return responseUser, nil
}

func (service *UserService) UpdateUser(c *gin.Context, input types.RequestUpdateUser) (types.ResponseUser, error) {
	currentUser := FindCurrentUser(c)

	if input.Name != "" {
		currentUser.Name = input.Name
	}
	if input.Nickname != "" {
		currentUser.Nickname = input.Nickname
	}
	if input.Email != "" {
		currentUser.Email = input.Email
	}

	updatedUser, err := service.UserRepository.Update(currentUser)
	if err != nil {
		return types.ResponseUser{}, err
	}

	responseUser := types.ResponseUser{
		ID:       updatedUser.ID,
		Name:     updatedUser.Name,
		Nickname: updatedUser.Nickname,
		Email:    updatedUser.Email,
	}
	return responseUser, nil
}

func (service *UserService) DeleteUser(c *gin.Context) error {
	currentUser := FindCurrentUser(c)

	deletePostsErr := service.PostRepository.DeleteAllByUserID(currentUser.ID)
	if deletePostsErr != nil {
		return deletePostsErr
	}

	err := service.UserRepository.Delete(currentUser)
	if err != nil {
		return err
	}

	return nil
}
