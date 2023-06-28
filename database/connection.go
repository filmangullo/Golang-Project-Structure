package database

// import (
// 	"fmt"
// 	"startupfundinggolang/config"
// 	"startupfundinggolang/models"

// 	"gorm.io/driver/mysql"w
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func ConnectToDatabase() (*gorm.DB, error) {
// 	dbConfig, err := config.Database()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load application configuration: %w", err)
// 	}

// 	var db *gorm.DB

// 	switch dbConfig.DBConnection {
// 	case "mysql":
// 		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 			dbConfig.DBUsername, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBDatabase)
// 		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to connect to MySQL database: %w", err)
// 		}
// 	case "postgres":
// 		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
// 			dbConfig.DBHost, dbConfig.DBUsername, dbConfig.DBPassword, dbConfig.DBDatabase, dbConfig.DBPort)
// 		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to connect to PostgreSQL database: %w", err)
// 		}
// 	default:
// 		return nil, fmt.Errorf("unsupported database connection: %s", dbConfig.DBConnection)
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get database instance: %w", err)
// 	}

// 	// Set connection pool settings
// 	sqlDB.SetMaxIdleConns(10)
// 	sqlDB.SetMaxOpenConns(100)

// 	// Perform database migration or other setup tasks if required
// 	db.AutoMigrate(&models.User{})
// 	// ...

// 	return db, nil
// }
