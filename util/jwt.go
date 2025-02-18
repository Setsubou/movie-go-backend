package util

import (
	"backend/errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func IsJWTValid(tokenInput string) error {
	token, err := jwt.Parse(tokenInput, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrBadRequest.SetMessage("unrecognized token")
		}

		return []byte("m3a7QDs8+ZCqLJvLOh/mfAiwX1MqUzWQbdHEZ6i5Tq8="), nil //TODO this returns secret key, move it to env var later
	})

	if err != nil {
		return errors.ErrBadRequest.SetMessage("unrecognized token")
	}

	alive := token.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().Unix(), true)
	if !alive {
		return errors.ErrUnauthorized.SetMessage("token is expired, please generate a new one")
	}

	return nil
}

func GenerateJWT(username string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "movie-backend",
		"jti": uuid.New(),
		"sub": username,
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	token_string, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", errors.ErrInternalError
	}

	return token_string, nil
}