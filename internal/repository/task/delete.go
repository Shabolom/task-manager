package task

import (
	"context"
	"fmt"
)

func (s *Storage) DeleteTask(ctx context.Context, taskID int64) (int64, error) {
	query := `
		UPDATE tasks
		SET
			deleted_at = NOW(),
			updated_at = NOW()
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	res, err := s.conn.Exec(ctx, query, taskID)
	if err != nil {
		return 0, fmt.Errorf("failed to delete task %d: %w", taskID, err)
	}

	return res.RowsAffected(), nil
}
