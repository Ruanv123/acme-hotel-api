package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ruanv123/acme-hotel-api/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: error loading .env file: %s\n", err)
	}

	// inicializando a conexão com o banco
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}
	return port
}
