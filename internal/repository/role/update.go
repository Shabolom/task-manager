package role

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) UpdateRole(ctx context.Context, updateReq dto.UpdateRoleRequest) (dto.Role, error) {
	query := `
		UPDATE roles
		SET
			name = $1,
			description = $2,
			updated_at = NOW()
		WHERE id = $3
		  AND deleted_at IS NULL
		RETURNING id, name, description, created_at, updated_at, deleted_at
	`

	var role dto.Role
	err := s.conn.QueryRow(ctx, query, updateReq.Name, updateReq.Description, updateReq.RoleID).Scan(
		&role.ID,
		&role.Name,
		&role.Description,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
	)
	if err != nil {
		return dto.Role{}, fmt.Errorf("failed to update role %d: %w", updateReq.RoleID, err)
	}

	return role, nil
}
