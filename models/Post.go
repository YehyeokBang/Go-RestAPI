package models

type Post struct {
	ID      uint `gorm:"primary_key"`
	Title   string
	Content string
	UserID  uint `gorm:"not null"`
	User    User `gorm:"foreignkey:UserID"`
}
