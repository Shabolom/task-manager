package refresh_token

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) GetRefreshTokenByID(ctx context.Context, tokenID int64) (*dto.RefreshToken, error) {
	query := `
		SELECT
			id,
			user_id,
			token,
			expires_at,
			revoked_at,
			created_at,
			updated_at,
			deleted_at
		FROM refresh_tokens
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	var token dto.RefreshToken

	err := s.conn.QueryRow(ctx, query, tokenID).Scan(
		&token.ID,
		&token.UserID,
		&token.Token,
		&token.ExpiresAt,
		&token.RevokedAt,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get refresh token by id %d: %w", tokenID, err)
	}

	return &token, nil
}
