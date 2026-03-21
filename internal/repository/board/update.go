package board

import (
	"context"
	"fmt"
	"time"
)

func (s *Storage) UpdateBoard(ctx context.Context, boardID int64, title, description string) (*time.Time, error) {
	query := `
		UPDATE boards
		SET
			title = $1,
			description = $2,
			updated_at = NOW()
		WHERE id = $3
		  AND deleted_at IS NULL
		RETURNING updated_at
	`

	var updatedAt time.Time

	err := s.conn.QueryRow(ctx, query, title, description, boardID).Scan(&updatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update board %d: %w", boardID, err)
	}

	return &updatedAt, nil
}
