package configuration

import (
	"fmt"
	"log"
	"strings"

	"github.com/joho/godotenv"
)

type ApplicationConfiguration struct {
	port         string
	host         string
	Release_mode string
	Secret       string
	Allowed_origins []string
}

func (a *ApplicationConfiguration) GetApplicationConnectionString() string {
	return fmt.Sprintf("%s:%s", a.host, a.port)
}

type DatabaseConfiguration struct {
	username      string
	password      string
	port          string
	host          string
	database_name string
	database_mode string
}

func (d *DatabaseConfiguration) GetDatabaseConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		d.username,
		d.password,
		d.host,
		d.port,
		d.database_name,
	)
}

type Configuration struct {
	ApplicationConfiguration ApplicationConfiguration
	DatabaseConfiguration    DatabaseConfiguration
}

func InitConfiguration() *Configuration {
	var config map[string]string
	config, err := godotenv.Read()

	if err != nil {
		log.Fatal("Missing .env")
	}

	allowedOriginsRaw := config["ALLOWED_ORIGINS"]
	allowedOrigins := []string{}

	for _, origin := range strings.Split(allowedOriginsRaw, ",") {
		trimmed := strings.TrimSpace(origin)
		if trimmed != "" {
			allowedOrigins = append(allowedOrigins, trimmed)
		}
	}
	
	return &Configuration{
		ApplicationConfiguration: ApplicationConfiguration{
			port:         config["APP_PORT"],
			host:         config["APP_HOST"],
			Release_mode: config["RELEASE_MODE"],
			Secret:       config["JWT_SECRET_KEY"],
			Allowed_origins: allowedOrigins,
		},
		DatabaseConfiguration: DatabaseConfiguration{
			username:      config["DB_USERNAME"],
			password:      config["DB_PASSWORD"],
			port:          config["DB_PORT"],
			host:          config["DB_HOST"],
			database_name: config["DB_NAME"],
			database_mode: config["DB_MODE"],
		},
	}
}
