package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"waiter-app/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Import SQLite dialect
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	// Run migrations (create tables if they don't exist)
	db.AutoMigrate(&models.Group{}, &models.Table{})  // Automatically migrate Group and Table models
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, waiter!")
	})

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	godotenv.Load(".env." + env + ".local")
	if env != "test" {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load()

	port := os.Getenv("PORT")

	fmt.Printf("Starting server on :%s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
