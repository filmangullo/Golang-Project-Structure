package main

import (
	"log"
	"your_project_name/config"
	"your_project_name/initializers"
	"your_project_name/utils/LoggerFunctions"
)

func init() {
	// Call the Database function from the initializers package
	initializers.Database()
}

func main() {
	// Load the application configuration
	appConfig, err := config.App()
	if err != nil {
		log.Fatalf("Failed to load application configuration: %v", err)
	}

	LoggerFunctions.Log.Info("Starting application")

	// Initialize the router
	r := initializers.Api()
	// Run the server
	r.Run(":" + appConfig.AppPort) // Change the port if needed
}
