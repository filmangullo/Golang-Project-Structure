package config

import (
	"log"

	"github.com/joho/godotenv"
)

/*
|--------------------------------------------------------------------------
| Environment Config
|--------------------------------------------------------------------------
|
| Here is where you can register data from .env for your application.
|
*/

type DatabaseConfig struct {
	DBHost       string
	DBConnection string
	DBPort       string
	DBDatabase   string
	DBUsername   string
	DBPassword   string
}

func Database() (DatabaseConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
		return DatabaseConfig{}, err
	}

	config := DatabaseConfig{
		DBHost:       environment("DB_HOST", "127.0.0.1"),
		DBConnection: environment("DB_CONNECTION", "mysql"),
		DBPort:       environment("DB_PORT", "3306"),
		DBDatabase:   environment("DB_DATABASE", "go"),
		DBUsername:   environment("DB_USER", "root"),
		DBPassword:   environment("DB_PASSWORD", ""),
	}

	return config, nil
}
