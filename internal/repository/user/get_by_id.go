package user

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) GetUserByID(ctx context.Context, userID int64) (*dto.User, error) {
	query := `
		SELECT
			id,
			role_id,
			email,
			password_hash,
			full_name,
			is_active,
			last_login_at,
			created_at,
			updated_at,
			deleted_at
		FROM users
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	var user dto.User

	err := s.conn.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.RoleID,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.IsActive,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id %d: %w", userID, err)
	}

	return &user, nil
}
