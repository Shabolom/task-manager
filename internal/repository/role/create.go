package role

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) CreateRole(ctx context.Context, role *dto.Role) error {
	query := `
		INSERT INTO roles (
			name,
			description
		)
		VALUES ($1, $2)
		RETURNING
			id,
			created_at,
			updated_at,
			deleted_at
	`

	err := s.conn.QueryRow(ctx, query,
		role.Name,
		role.Description,
	).Scan(
		&role.ID,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create role: %w", err)
	}

	return nil
}
