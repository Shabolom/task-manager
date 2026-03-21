package refresh_token

import (
	"context"
	"fmt"
	"time"
)

func (s *Storage) UpdateRefreshToken(
	ctx context.Context,
	tokenID int64,
	expiresAt time.Time,
	revokedAt *time.Time,
) (*time.Time, error) {
	query := `
		UPDATE refresh_tokens
		SET
			expires_at = $1,
			revoked_at = $2,
			updated_at = NOW()
		WHERE id = $3
		  AND deleted_at IS NULL
		RETURNING updated_at
	`

	var updatedAt time.Time

	err := s.conn.QueryRow(ctx, query, expiresAt, revokedAt, tokenID).Scan(&updatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update refresh token %d: %w", tokenID, err)
	}

	return &updatedAt, nil
}
