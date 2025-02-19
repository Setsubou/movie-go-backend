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
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin"},
		AllowCredentials: true,
	}))

	postgres := db.NewPostgresDb(config.DatabaseConfiguration.GetDatabaseConnectionString())
	jwt_secret := config.ApplicationConfiguration.Secret
	jwt_middleware := jwt.JWT(jwt_secret)

	router.GET("/health-check/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Auth endpoints
	auth_controller := api.NewAuthController(postgres, jwt_secret)
	router.POST("/auth/", auth_controller.VerifyUserLogin)
	router.POST("/auth/verify-token/", jwt_middleware, auth_controller.VerifyToken)

	apiv1 := router.Group("/api/v1/")
	apiv1.Use(jwt_middleware)

	// Movie endpoints
	movie_controller := api.NewMovieController(postgres)

	apiv1.POST("/movie/", movie_controller.InsertNewMovie)
	apiv1.GET("/movie/:id/", movie_controller.GetMovieById)
	apiv1.DELETE("/movie/:id/", movie_controller.DeleteMovieById)
	apiv1.GET("/movie/byPublisher/:id/", movie_controller.GetMoviesByPublisherId)

	apiv1.GET("/movies/", movie_controller.GetListAllMovies)

	// Publisher endpoints
	publisher_controller := api.NewPublisherController(postgres)
	apiv1.GET("/publishers/name/", publisher_controller.GetListAllPublishersName)

	// Genre endpoints
	genre_controller := api.NewGenreController(postgres)
	apiv1.GET("/genres/", genre_controller.GetAllGenres)

	return router
}
