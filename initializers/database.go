package initializers

import (
	"fmt"
	"startupfundinggolang/database"
)

func Database() {
	// Call the ConnectToDatabase function from the database package
	db, err := database.ConnectToDatabase()
	if err != nil {
		// Handle the error if connection fails
		fmt.Printf("Failed to connect to the database: %s\n", err)
		// You can choose to exit the program here if necessary
		return
	}

	// Close the database connection when the program exits
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Failed to get database instance: %s\n", err)
		// Handle the error if necessary
		return
	}
	defer sqlDB.Close()
}
