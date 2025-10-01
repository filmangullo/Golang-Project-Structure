package main

import (
	"log"
	"your_project_name/config"
	"your_project_name/initializers"
)

func init() {
	// Call the Root function from the initializers package
	initializers.Root()
	// Call the Database function from the initializers package
	initializers.Database()
}

func main() {
	// Load the application configuration
	appConfig, err := config.App()
	if err != nil {
		log.Fatalf("Failed to load application configuration: %v", err)
	}

	// Initialize the router
	r := initializers.Api()
	// Run the server
	r.Run(":" + appConfig.AppPort) // Change the port if needed
}
