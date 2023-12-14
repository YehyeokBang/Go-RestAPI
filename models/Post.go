package models

type Post struct {
	ID      uint `gorm:"primary_key"`
	Title   string
	Content string
	UserID  uint `gorm:"not null"`          // 외래 키 추가
	User    User `gorm:"foreignkey:UserID"` // 외래 키 관계 설정
}
