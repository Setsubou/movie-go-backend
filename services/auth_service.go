package services

import (
	"backend/errors"
	"backend/repository"
	"backend/util"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository repository.AuthService
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

func NewAuthService(repository repository.AuthService) *AuthService {
	return &AuthService{
		repository: repository,
	}
}
