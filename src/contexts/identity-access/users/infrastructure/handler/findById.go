package handler

import (
	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/application"
	sharedInfrastructure "github.com/ezflow-me/identity-management-service/src/contexts/shared/infrastructure"
	"github.com/gofiber/fiber/v2"
)

func FindByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		user, err := application.FindByID(c.Context(), id)
		if err != nil {
			return sharedInfrastructure.NewErrResponse(c, err, fiber.StatusInternalServerError)
		}

		return sharedInfrastructure.NewSuccessResponse(c, user.ToPrimitives())
	}
}
