package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"type:varchar(6);not null"`
}

//not for migration

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
