package router

import (
	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/infrastructure/handler"
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Post("/register", handler.Register())
	router.Get("/:id", handler.FindByID())
	router.Delete("/:id", handler.Delete())
}
