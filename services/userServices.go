package services

import (
	"bloggers/middleware"
	"bloggers/models"
	"bloggers/repo"
	"bloggers/utils"
	"errors"
)

type UserServices struct {
	Repo repo.UserRepo
}

func (s *UserServices) RegisterUser(user *models.User) error {
	//check if user exists
	existingUser, err := s.Repo.GetUserByEmail(user.Email)

	if err != nil {
		return err //DB error
	}

	if existingUser != nil {
		return errors.New("this email already has an account")
	}

	//hash password
	hashpass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashpass

	//call create method
	err = s.Repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil

}

func (s *UserServices) LogInUser(request models.LoginRequest) (string, error) {
	user, err := s.Repo.GetUserByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, request.Password)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
