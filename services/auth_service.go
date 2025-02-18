package services

import (
	"backend/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

func (s *AuthService) VerifyUserLogin(username string, password string) (string, error) {
	hashed_password := s.repository.GetOneUserByUsernameAndPasswordHash(username) //TODO change this to just by Username

	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))

	//TODO move token generation into util, also move key to env
	if err == nil { //TODO reverse the if to make early return
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": "movie-backend",
			"jti": uuid.New(),
			"sub": username,
			"nbf": time.Now().Unix(),
			"iat": time.Now().Unix(),
			"exp": time.Now().Add(time.Hour).Unix(),
		})
		token_string, _ := token.SignedString([]byte("m3a7QDs8+ZCqLJvLOh/mfAiwX1MqUzWQbdHEZ6i5Tq8=")) //TODO move to env var

		return token_string, nil
	}

	return "", errors.New("Credential not found")
}

