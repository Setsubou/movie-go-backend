package routes

import (
	db "backend/db/postgres_db"
	"fmt"
	"os"
	"context"
	
	"backend/routes/api"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRouter() *gin.Engine {
	// Init DB selection here
	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:password@localhost:5432/movie") //TODO replace this with env var later
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	
	router := gin.New()

	router.GET("/health-check/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "OK",
		})
	})

	movie_controller := api.NewMovieController(db.NewPostgresDb(dbpool))
	router.GET("movie/:id", movie_controller.GetMovieById)

	return router
}
