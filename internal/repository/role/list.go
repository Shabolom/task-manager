package role

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) ListRoles(ctx context.Context, pagination dto.Pagination) ([]dto.Role, error) {
	query := `
		SELECT
			id,
			name,
			description,
			created_at,
			updated_at,
			deleted_at
		FROM roles
		WHERE deleted_at IS NULL
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := s.conn.Query(ctx, query, pagination.Limit, pagination.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list roles: %w", err)
	}
	defer rows.Close()

	roles := make([]dto.Role, 0)

	for rows.Next() {
		var role dto.Role

		err = rows.Scan(
			&role.ID,
			&role.Name,
			&role.Description,
			&role.CreatedAt,
			&role.UpdatedAt,
			&role.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan role: %w", err)
		}

		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while iterating roles: %w", err)
	}

	return roles, nil
}
