package statuses

import (
	"context"
	"fmt"
)

func (s *Storage) DeleteStatus(ctx context.Context, statusID int64) (int64, error) {
	query := `
		UPDATE statuses
		SET
			deleted_at = NOW(),
			updated_at = NOW()
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	res, err := s.conn.Exec(ctx, query, statusID)
	if err != nil {
		return 0, fmt.Errorf("failed to delete status %d: %w", statusID, err)
	}

	return res.RowsAffected(), nil
}
