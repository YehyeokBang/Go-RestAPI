package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}