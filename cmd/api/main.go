package main

import (
	"log"
	"os"

	"github.com/ezflow-me/identity-management-service/server"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	app := server.Setup()

	if os.Getenv("STAGE_STATUS") == "dev" {
		if err := server.StartServer(app); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	} else {
		server.StartServerWithGracefulShutdown(app)
	}
}
