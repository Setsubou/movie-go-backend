package jwt

import (
	"backend/util"
	"backend/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		authToken := strings.TrimPrefix(authHeader, "Bearer ")

		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "missing bearer token.",
			})

			c.Abort()
			return
		}

		err := util.IsJWTValid(authToken)

		if err != nil {
			if internalErr, ok := err.(*errors.InternalError); ok {
				c.JSON(internalErr.Code, gin.H{"error": internalErr.Message})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}

			c.Abort()
			return
		}

		c.Next()
	}
}
