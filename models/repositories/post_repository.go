package repositories

import (
	"example/board/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		DB: db,
	}
}

func (repo *PostRepository) Create(post *models.Post) (*models.Post, error) {
	if err := repo.DB.Create(post).Error; err != nil {
		return nil, err
	}
	if err := repo.DB.Model(post).Preload("User").First(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (repo *PostRepository) FindAll() ([]models.Post, error) {
	var posts []models.Post
	result := repo.DB.Preload("User").Find(&posts)
	return posts, result.Error
}

func (repo *PostRepository) FindByID(id uint) (models.Post, error) {
	var post models.Post
	result := repo.DB.Preload("User").Where("id = ?", id).First(&post)
	return post, result.Error
}

func (repo *PostRepository) Update(post *models.Post) (*models.Post, error) {
	if err := repo.DB.Save(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *PostRepository) Delete(post *models.Post) error {
	return repo.DB.Delete(post).Error
}

func (repo *PostRepository) DeleteAllByUserID(userID uint) error {
	return repo.DB.Where("user_id = ?", userID).Delete(&models.Post{}).Error
}
