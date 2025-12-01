package server

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ezflow-me/identity-management-service/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Setup() *fiber.App {
	config := Config()
	app := fiber.New(config)

	middleware.AddFiberMiddleware(app)

	SetupRoutes(app)

	return app
}

func StartServerWithGracefulShutdown(app *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS interrupt signal
		<-sigint

		// We received an interrupt signal, shut down.
		if err := app.Shutdown(); err != nil {
			log.Error("Server Shutdown Failed: %v", err)
		}
		close(idleConnsClosed)
	}()

	fiberConnURL := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))

	if err := app.Listen(fiberConnURL); err != nil {
		log.Error("Server Start Failed: %v", err)
	}

	<-idleConnsClosed
}

func StartServer(app *fiber.App) error {
	fiberConnURL := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))

	return app.Listen(fiberConnURL)
}
