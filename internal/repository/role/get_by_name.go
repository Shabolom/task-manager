package role

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) GetRoleByName(ctx context.Context, roleName string) (*dto.Role, error) {

	query := `
		SELECT
			id,
			name,
			description,
			created_at,
			updated_at,
			deleted_at
		FROM roles
		WHERE name = $1
		AND deleted_at IS NULL
	`

	var role dto.Role

	err := s.conn.QueryRow(ctx, query, roleName).Scan(
		&role.ID,
		&role.Name,
		&role.Description,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get role by name %d: %w", roleName, err)
	}

	return &role, nil
}
