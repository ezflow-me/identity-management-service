package postgre

import (
	"context"

	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/domain"
)

func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return domain.NewErrCannotDeleteUser(id, err.Error())
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.NewErrUserNotFound(id)
	}

	return nil
}
