package statuses

import (
	"context"
	"fmt"
	"time"
)

func (s *Storage) UpdateStatus(
	ctx context.Context,
	statusID int64,
	name string,
	color string,
) (*time.Time, error) {
	query := `
		UPDATE statuses
		SET
			name = $1,
			color = $2,
			updated_at = NOW()
		WHERE id = $3
		  AND deleted_at IS NULL
		RETURNING updated_at
	`

	var updatedAt time.Time

	err := s.conn.QueryRow(ctx, query, name, color, statusID).Scan(&updatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update status %d: %w", statusID, err)
	}

	return &updatedAt, nil
}
