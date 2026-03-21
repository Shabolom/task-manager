package user

import (
	"context"
	"fmt"
)

func (s *Storage) DeleteUser(ctx context.Context, userID int64) (int64, error) {
	query := `
		UPDATE users
		SET
			deleted_at = NOW(),
			updated_at = NOW()
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	res, err := s.conn.Exec(ctx, query, userID)
	if err != nil {
		return 0, fmt.Errorf("failed to delete user %d: %w", userID, err)
	}

	return res.RowsAffected(), nil
}
