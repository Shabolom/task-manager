package statuses

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) GetStatusByID(ctx context.Context, statusID int64) (*dto.Status, error) {
	query := `
		SELECT
			id,
			board_id,
			name,
			color,
			created_at,
			updated_at,
			deleted_at
		FROM statuses
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	var status dto.Status

	err := s.conn.QueryRow(ctx, query, statusID).Scan(
		&status.ID,
		&status.BoardID,
		&status.Name,
		&status.Color,
		&status.CreatedAt,
		&status.UpdatedAt,
		&status.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get status by id %d: %w", statusID, err)
	}

	return &status, nil
}
