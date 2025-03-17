package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite driver
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	dbName := os.Getenv("DB")
	DB, err = gorm.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to the database successfully.")
}

// CloseDB closes the database connection
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Failed to close the database connection: %v", err)
	}
	log.Println("Database connection closed.")
}
