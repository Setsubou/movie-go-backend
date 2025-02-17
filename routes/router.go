package routes

import (
	"backend/db"
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

	movie_controller := api.NewMovieController(db.NewMockDb())
	router.GET("movie/:id", movie_controller.GetMovieById)

	return router
}
