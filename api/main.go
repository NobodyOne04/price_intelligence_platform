package main

import (
	"log"
	"os"
	"context"

	"github.com/joho/godotenv"
	"api/db"
	"api/router"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file was found")
	}

	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("Cannot connect to the DB")
	}
	defer conn.Close(context.Background())

	connWrapper := &db.ConnWrapper(Conn: conn)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("API is launced on port: %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server instance failed: %v", err)
	}
}
