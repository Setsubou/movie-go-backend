package services

import (
	"backend/repository"
	"backend/util"
	"backend/errors"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(repository repository.AuthRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

type AuthService struct {
	repository repository.AuthRepository
}

func (s *AuthService) VerifyUserLogin(username string, password string, secret string) (string, error) {
	hashed_password := s.repository.GetOneUserByUsername(username)
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))

	if err != nil {
		return "", errors.ErrUnauthorized
	}

	token, err := util.GenerateJWT(username, secret)

	if err != nil {
		return "", err
	}

	return token, nil

}
