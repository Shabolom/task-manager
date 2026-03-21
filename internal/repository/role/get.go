package role

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) GetRoleByID(ctx context.Context, roleID int64) (*dto.Role, error) {

	query := `
		SELECT
			id,
			name,
			description,
			created_at,
			updated_at,
			deleted_at
		FROM roles
		WHERE id = $1
		AND deleted_at IS NULL
	`

	var role dto.Role

	err := s.conn.QueryRow(ctx, query, roleID).Scan(
		&role.ID,
		&role.Name,
		&role.Description,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get role by id %d: %w", roleID, err)
	}

	return &role, nil
}
