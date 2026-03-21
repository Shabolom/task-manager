package role

import (
	"context"
	"fmt"
)

func (s *Storage) DeleteRole(ctx context.Context, roleID int64) (int64, error) {
	query := `
		DELETE FROM roles
		WHERE id = $1
	`

	res, err := s.conn.Exec(ctx, query, roleID)
	if err != nil {
		return 0, fmt.Errorf("failed to delete role %d: %w", roleID, err)
	}

	return res.RowsAffected(), nil
}
