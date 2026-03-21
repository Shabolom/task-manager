package user

import (
	"context"
	"fmt"
)

func (s *Storage) UpdateUserPassword(ctx context.Context, userID int64, passwordHash []byte) (int64, error) {
	query := `
		UPDATE users
		SET
			password_hash = $1,
			updated_at = NOW()
		WHERE id = $2
		  AND deleted_at IS NULL
	`

	res, err := s.conn.Exec(ctx, query, passwordHash, userID)
	if err != nil {
		return 0, fmt.Errorf("failed to update password for user %d: %w", userID, err)
	}

	return res.RowsAffected(), nil
}
