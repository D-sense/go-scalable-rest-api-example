package config

import (
	"github.com/d-sense/go-scalable-rest-api-example/config/flags"
	"os"
)

// LoadEnvFile LoadConfig loads config from files
func LoadEnvFile() (*flags.Configuration, error) {
	var appConf flags.Configuration

	appConf.Postgres.Database = os.Getenv("DB_NAME")
	appConf.Postgres.Driver = os.Getenv("DB_DRIVER")
	appConf.Postgres.Username = os.Getenv("DB_USERNAME")
	appConf.Postgres.Password = os.Getenv("DB_PASSWORD")
	appConf.Postgres.Host = os.Getenv("DB_HOST")
	appConf.Postgres.Port = os.Getenv("DB_PORT")
	appConf.Address = os.Getenv("ADDRESS")

	return &appConf, nil
}
