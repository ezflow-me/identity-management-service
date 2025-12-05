package application

import (
	"context"

	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/domain"
	sharedInfrastructure "github.com/ezflow-me/identity-management-service/src/contexts/shared/infrastructure"
)

func Register(ctx context.Context, user *domain.User) error {
	err := sharedInfrastructure.Validate.Struct(user)
	if err != nil {
		return domain.NewErrInvalidUserData(err.Error())
	}

	if err := Repository.Save(ctx, user); err != nil {
		return domain.NewErrCannotSaveUser(err.Error())
	}

	return nil
}
