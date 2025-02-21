package api

import (
	"backend/errors"
	"backend/model"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAuthController(auth_service *services.AuthService, secret string) *Auth_controller {
	return &Auth_controller{
		auth_service: auth_service,
		secret:  secret,
	}
}

type Auth_controller struct {
	auth_service *services.AuthService
	secret  string
}

func (ac *Auth_controller) VerifyUserLogin(c *gin.Context) {
	var login_credentials model.Auth

	if err := c.ShouldBindBodyWithJSON(&login_credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required data are either missing or malformed."})
		return
	}

	token, err := ac.auth_service.VerifyUserLogin(login_credentials.Username, login_credentials.Password, ac.secret)

	if err != nil {
		if internalErr, ok := err.(*errors.InternalError); ok {
			c.JSON(internalErr.Code, gin.H{"error": internalErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.SetCookie("access_token", token, 3600, "/", "localhost", false, true)
	c.SetSameSite(http.SameSiteNoneMode)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (ac *Auth_controller) VerifyToken(c *gin.Context) {
	// Verification already happened in JWT middleware, if token is valid this function is run

	c.JSON(http.StatusOK, gin.H{
		"message": "token is still valid",
	})
}

func (ac *Auth_controller) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H {
		"message": "logged out successfully",
	})
}
