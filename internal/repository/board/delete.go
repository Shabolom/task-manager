package board

import (
	"context"
	"fmt"
)

func (s *Storage) DeleteBoard(ctx context.Context, boardID int64) (int64, error) {
	query := `
		UPDATE boards
		SET
			deleted_at = NOW(),
			updated_at = NOW()
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	res, err := s.conn.Exec(ctx, query, boardID)
	if err != nil {
		return 0, fmt.Errorf("failed to delete board %d: %w", boardID, err)
	}

	return res.RowsAffected(), nil
}
