package db

import (
	"log"
	"waiter-app/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite driver
)

// Migrate handles the DB migration (creating the tables).
func Migrate() {
	DB.AutoMigrate(&models.Group{}, &models.Table{})
	log.Println("Database migration completed.")
}
