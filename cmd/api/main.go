package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"github.com/ezflow-me/identity-management-service/server"
	_ "github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/application"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	app := server.Setup()

	if os.Getenv("STAGE_STATUS") == "dev" {
		if err := server.StartServer(app); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	} else {
		server.StartServerWithGracefulShutdown(app)
	}
}
