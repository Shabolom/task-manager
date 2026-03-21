package user

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
	"time"
)

func (s *Storage) UpdateUser(
	ctx context.Context,
	userID int64,
	request *dto.UpdateUserRequest,
) (*time.Time, error) {
	query := `
		UPDATE users
		SET
			role_id = COALESCE($1, role_id),
			email = COALESCE($2, email),
			full_name = COALESCE($3, full_name),
			is_active = COALESCE($4, is_active),
			last_login_at = COALESCE($5, last_login_at),
			updated_at = NOW()
		WHERE id = $6
		  AND deleted_at IS NULL
		RETURNING updated_at
	`

	var updatedAt time.Time

	err := s.conn.QueryRow(
		ctx,
		query,
		request.RoleID,
		request.Email,
		request.FullName,
		request.IsActive,
		request.LastLoginAt,
		userID,
	).Scan(&updatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update user %d: %w", userID, err)
	}

	return &updatedAt, nil
}
