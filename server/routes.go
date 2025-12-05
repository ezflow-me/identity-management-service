package server

import (
	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/infrastructure/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/metrics", monitor.New())

	user := app.Group("/user")
	router.Router(user)
}
