package repositories

import (
	"example/board/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := repo.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	result := repo.DB.Where("id = ?", id).First(&user)
	return user, result.Error
}

func (repo *UserRepository) Update(user *models.User) (*models.User, error) {
	if err := repo.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Delete(user *models.User) error {
	return repo.DB.Delete(user).Error
}
