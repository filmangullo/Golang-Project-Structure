package config

import (
	"os"
)

/*
|--------------------------------------------------------------------------
| Environment Config
|--------------------------------------------------------------------------
|
| Here is where you can register data from .env for your application.
|
*/

func environment(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
