package user

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) ListUsers(ctx context.Context) ([]dto.User, error) {
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
		WHERE deleted_at IS NULL
		ORDER BY id DESC
	`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	users := make([]dto.User, 0)

	for rows.Next() {
		var user dto.User

		err = rows.Scan(
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
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while iterating users: %w", err)
	}

	return users, nil
}
