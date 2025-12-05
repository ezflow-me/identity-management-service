package handler

import (
	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/application"
	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/domain"
	sharedInfrastructure "github.com/ezflow-me/identity-management-service/src/contexts/shared/infrastructure"
	"github.com/gofiber/fiber/v2"
)

func Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		command := new(UserCommand)
		if err := c.BodyParser(command); err != nil {
			return sharedInfrastructure.NewErrResponse(c, err, fiber.StatusBadRequest)
		}

		user, err := domain.NewUser(command.ID, command.Name, command.Email, command.CreatedAt)
		if err != nil {
			return sharedInfrastructure.NewErrResponse(c, err, fiber.StatusBadRequest)
		}

		err = application.Register(c.Context(), user)
		if err != nil {
			return sharedInfrastructure.NewErrResponse(c, err, fiber.StatusInternalServerError)
		}

		return sharedInfrastructure.NewSuccessResponse(c, user.ToPrimitives())
	}
}
