package services

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"safpass-api/models"
	"safpass-api/repository"
	"safpass-api/utils"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) CreateUser(user *models.User) error {
	user.ID = uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthService) AuthenticateUser(username string, password string) (*models.User, error) {
	user, err := repository.AuthenticateUser(username)
	if err != nil {
		return nil, err
	}

	// Compare password
	err = utils.ComparePassword([]byte(user.Password), password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
