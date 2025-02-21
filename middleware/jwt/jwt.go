package jwt

import (
	"backend/errors"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWT(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("access_token")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "missing access token.",
			})

			c.Abort()
			return
		}

		if err := util.IsJWTValid(cookie, secret); err != nil {
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
