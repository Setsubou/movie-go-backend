package routes

import (
	db "backend/db/postgres_db"
	"backend/middleware/jwt"
	"context"
	"fmt"
	"os"

	"backend/routes/api"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRouter(database_connection_url string) *gin.Engine { //TODO make it so this accept config instead of db connection, more flexible
	// I also need to pass config to auth so it can access env key later
	dbpool, err := pgxpool.New(context.Background(), database_connection_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health-check/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "OK",
		})
	})

	auth_controller := api.NewAuthController(db.NewPostgresDb(dbpool))
	router.POST("/auth/", auth_controller.VerifyUserLogin)

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT()) 
	
	movie_controller := api.NewMovieController(db.NewPostgresDb(dbpool))
	apiv1.GET("/movie/:id", movie_controller.GetMovieById)
	apiv1.GET("/movies/", movie_controller.GetListAllMovies)

	return router
}
