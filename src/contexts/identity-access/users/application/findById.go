package application

import (
	"context"

	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/domain"
)

func FindByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := Repository.FindByID(ctx, id)
	if err != nil {
		return nil, domain.NewErrCannotFindUserByID(id, err.Error())
	}

	return user, nil
}
