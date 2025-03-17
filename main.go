package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"waiter-app/db"
	"waiter-app/routes"
)

func main() {
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

	db.InitDB()
	defer db.CloseDB()
	db.Migrate()

	routes.RegisterRoutes()

	fmt.Printf("Starting server on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
