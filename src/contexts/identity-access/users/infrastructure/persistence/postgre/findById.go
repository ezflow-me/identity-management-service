package postgre

import (
	"context"

	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/domain"
)

func (r *Repository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	query := `
		SELECT id, name, email, time_zone, display_language, created_at
		FROM users
		WHERE id = $1
	`

	user := &domain.User{}

	err := r.db.QueryRow(ctx, query, id).Scan(
		user.PtrsForScan()...,
	)

	if err != nil {
		return nil, domain.NewErrCannotFindUserByID(id, err.Error())
	}

	return user, nil
}
