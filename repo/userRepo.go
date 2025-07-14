package repo

import (
	"bloggers/database"
	"bloggers/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
}

type UserRepo struct {
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil //no user found but not an error
	}
	if err != nil {
		return &models.User{}, err //error
	}

	return &user, nil //user found
}

func (r *UserRepo) CreateUser(user *models.User) error {
	err := database.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
