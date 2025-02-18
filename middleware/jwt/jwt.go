package jwt

import (
	"backend/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		authToken := strings.TrimPrefix(authHeader, "Bearer ")

		if authToken == "" { //TODO check if token is valid and check if token has expired or not
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "missing bearer token.",
			})

			c.Abort()
			return
		}

		valid, err := util.IsJWTValid(authToken)
		if !valid || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
