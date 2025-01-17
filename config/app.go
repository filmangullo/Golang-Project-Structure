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

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func App() (AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
		return AppConfig{}, err
	}

	config := AppConfig{
		AppName: environment("APP_NAME", "Golang Project"),
		AppEnv:  environment("APP_ENV", "development"),
		AppPort: environment("APP_PORT", "8000"),
	}

	return config, nil
}
