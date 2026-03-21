package user

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) CreateUser(ctx context.Context, user *dto.User) error {
	query := `
		INSERT INTO users (
			role_id,
			email,
			password_hash,
			full_name,
			is_active,
			last_login_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING
			id,
			created_at,
			updated_at,
			deleted_at
	`

	err := s.conn.QueryRow(ctx, query,
		user.RoleID,
		user.Email,
		user.PasswordHash,
		user.FullName,
		user.IsActive,
		user.LastLoginAt,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
