package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content" gorm:"not null"`
	PostID  uint   `json:"post_id" gorm:"not null"`
	UserID  uint   `json:"user_id" gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"`
}

type CommentContent struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}
