package main

import (
	"backend/configuration"
	"backend/routes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	configuration := configuration.InitConfiguration()

	gin.SetMode(configuration.ApplicationConfiguration.Release_mode)
	router := routes.InitRouter(configuration.DatabaseConfiguration.GetDatabaseConnectionString())

	server := &http.Server{
		Addr:    configuration.ApplicationConfiguration.GetApplicationConnectionString(),
		Handler: router,
	}

	fmt.Printf("Listening at %s", server.Addr)

	server.ListenAndServe()
}
