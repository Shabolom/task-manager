package refresh_token

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) ListRefreshTokensByUser(ctx context.Context, userID int64) ([]dto.RefreshToken, error) {
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
		WHERE user_id = $1
		  AND deleted_at IS NULL
		ORDER BY created_at DESC
	`

	rows, err := s.conn.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list refresh tokens by user %d: %w", userID, err)
	}
	defer rows.Close()

	tokens := make([]dto.RefreshToken, 0)

	for rows.Next() {
		var token dto.RefreshToken

		err = rows.Scan(
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
			return nil, fmt.Errorf("failed to scan refresh token: %w", err)
		}

		tokens = append(tokens, token)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while iterating refresh tokens: %w", err)
	}

	return tokens, nil
}
