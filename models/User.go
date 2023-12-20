package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func GetUserByEmail(email string) (User, error) {
	var user User
	result := DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
