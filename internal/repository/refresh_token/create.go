package refresh_token

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) CreateRefreshToken(ctx context.Context, token *dto.RefreshToken) error {
	query := `
		INSERT INTO refresh_tokens (
			user_id,
			token,
			expires_at,
			revoked_at
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id,
			created_at,
			updated_at,
			deleted_at
	`

	err := s.conn.QueryRow(ctx, query,
		token.UserID,
		token.Token,
		token.ExpiresAt,
		token.RevokedAt,
	).Scan(
		&token.ID,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create refresh token: %w", err)
	}

	return nil
}
