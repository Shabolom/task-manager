package refresh_token

import (
	"context"
	"fmt"
)

func (s *Storage) DeleteRefreshToken(ctx context.Context, tokenID int64) (int64, error) {
	query := `
		UPDATE refresh_tokens
		SET
			deleted_at = NOW(),
			updated_at = NOW()
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	res, err := s.conn.Exec(ctx, query, tokenID)
	if err != nil {
		return 0, fmt.Errorf("failed to delete refresh token %d: %w", tokenID, err)
	}

	return res.RowsAffected(), nil
}
