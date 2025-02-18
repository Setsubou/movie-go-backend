package api

import (
	"backend/errors"
	"backend/model"
	"backend/repository"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAuthController(repository repository.AuthRepository, secret string) *Auth_controller {
	return &Auth_controller{
		repository: repository,
		secret: secret,
	}
}

type Auth_controller struct {
	repository repository.AuthRepository
	secret string
}

func (ac *Auth_controller) VerifyUserLogin(c *gin.Context) {
	var login_credentials model.Auth

	if err := c.ShouldBindBodyWithJSON(&login_credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required data are either missing or malformed."})
		return
	}

	token, err := services.NewAuthService(ac.repository).VerifyUserLogin(login_credentials.Username, login_credentials.Password, ac.secret)
	
	if err != nil {
		if internalErr, ok := err.(*errors.InternalError); ok {
            c.JSON(internalErr.Code, gin.H{"error": internalErr.Message})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        }
        return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	}) //Temp status for now, Use cookies instead
}
