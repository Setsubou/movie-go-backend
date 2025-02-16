package routes

import (
	"backend/routes/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/health-check/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "OK",
		})
	})

	router.GET("movie/:id", api.GetMovie)

	return router
}