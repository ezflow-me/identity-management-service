package postgre

import (
	"context"

	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/domain"
)

func (r *Repository) Save(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (id, name, email, created_at, display_language, time_zone)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET
			name = EXCLUDED.name,
			email = EXCLUDED.email,
			created_at = EXCLUDED.created_at,
	`

	_, err := r.db.Exec(
		ctx,
		query,
		user.ID(),
		user.Name(),
		user.Email(),
		user.CreatedAt(),
		user.DisplayLanguage(),
		user.TimeZone(),
	)

	if err != nil {
		return domain.NewErrCannotSaveUser(err.Error())
	}

	return nil
}
