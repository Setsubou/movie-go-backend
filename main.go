package main

import (
	"backend/routes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug") //TODO change this to env file later

	router := routes.InitRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Listening at %s", server.Addr)

	server.ListenAndServe()
}
