package application

import (
	"context"

	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/domain"
)

func Delete(ctx context.Context, id string) error {
	err := Repository.Delete(ctx, id)
	if err != nil {
		return domain.NewErrCannotDeleteUser(id, err.Error())
	}

	return nil
}
