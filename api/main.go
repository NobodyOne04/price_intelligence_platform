package main

import (
	"context"
	"log"
	"os"

	"api/db"
	"api/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file was found")
	}

	pool, err := db.Connect(context.Background())
	if err != nil {
		log.Fatalf("Cannot connect to the DB: %v", err)
	}
	defer pool.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("API is launched on port: %s", port)

	if err := router.SetupRouter(pool).Run(":" + port); err != nil {
		log.Fatalf("Server instance failed: %v", err)
	}
}
