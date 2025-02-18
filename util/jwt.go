package util

import (
	"errors"
	"fmt"
	"time"

	//"time"

	"github.com/golang-jwt/jwt"
)

func IsJWTValid(tokenInput string) (bool, error) {

	fmt.Println("Checking token validity")

	token, err := jwt.Parse(tokenInput, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong signing signature")
		}

		return []byte("m3a7QDs8+ZCqLJvLOh/mfAiwX1MqUzWQbdHEZ6i5Tq8="), nil //TODO this returns secret key, move it to env var later
	})

	if err != nil {
		return false, errors.New("malformed token")
	}

	alive := token.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().Unix(), true)
	if !alive {
		return false, errors.New("token is expired, please generate a new one")
	}

	return true, nil
}