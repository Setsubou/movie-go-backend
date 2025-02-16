package main

import (
	"backend/routes"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:password@localhost:5432/movie") //TODO replace this with env var later
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	gin.SetMode("debug") //TODO change this to env file later

	router := routes.InitRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Listening at %s", server.Addr)

	server.ListenAndServe()

	dbpool.Close()
}
