package routes

import (
	"backend/configuration"
	db "backend/db/postgres_db"
	"backend/middleware/jwt"
	"backend/routes/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(config *configuration.Configuration) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"PUT", "GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin"},
		AllowCredentials: true,
	}))

	postgres := db.NewPostgresDb(config.DatabaseConfiguration.GetDatabaseConnectionString())

	router.GET("/health-check/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "OK",
		})
	})

	auth_controller := api.NewAuthController(postgres, config.ApplicationConfiguration.Secret)
	router.POST("/auth/", auth_controller.VerifyUserLogin)

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT()) 
	
	movie_controller := api.NewMovieController(postgres)
	apiv1.GET("/movie/:id", movie_controller.GetMovieById)
	apiv1.GET("/movies/", movie_controller.GetListAllMovies)

	return router
}
