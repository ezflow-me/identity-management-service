package infrastructure

import (
	"github.com/gofiber/fiber/v2"
)

func NewErrResponse(c *fiber.Ctx, err error, status int) error {
	return c.Status(status).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}

func NewSuccessResponse(c *fiber.Ctx, data ...any) error {
	return c.Status(fiber.StatusOK).JSON(data)
}
