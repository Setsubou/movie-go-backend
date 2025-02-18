package main

import (
	"backend/configuration"
	"backend/routes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configuration.InitConfiguration()

	gin.SetMode(config.ApplicationConfiguration.Release_mode)
	router := routes.InitRouter(config)

	server := &http.Server{
		Addr:    config.ApplicationConfiguration.GetApplicationConnectionString(),
		Handler: router,
	}

	fmt.Printf("\n listening at %s \n", server.Addr)
	server.ListenAndServe()
}
