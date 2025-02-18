package api

import (
	"backend/model"
	"backend/repository"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAuthController(repository repository.AuthRepository) *Auth_controller {
	return &Auth_controller{
		repository: repository,
	}
}

type Auth_controller struct {
	repository repository.AuthRepository
}

func (ac *Auth_controller) VerifyUserLogin(c *gin.Context) {
	var login_credentials model.Auth

	if err := c.ShouldBindBodyWithJSON(&login_credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required data are either missing or malformed."})
		return
	}

	token, err := services.NewAuthService(ac.repository).VerifyUserLogin(login_credentials.Username, login_credentials.Password)
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong credentials",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.JSON(http.StatusOK, gin.H {
		"token": token,
	}) //Temp status for now
}
