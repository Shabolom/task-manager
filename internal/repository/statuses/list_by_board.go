package statuses

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) ListStatusesByBoard(ctx context.Context, boardID int64) ([]dto.Status, error) {
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
		WHERE board_id = $1
		  AND deleted_at IS NULL
		ORDER BY position ASC, id ASC
	`

	rows, err := s.conn.Query(ctx, query, boardID)
	if err != nil {
		return nil, fmt.Errorf("failed to list statuses by board %d: %w", boardID, err)
	}
	defer rows.Close()

	statuses := make([]dto.Status, 0)

	for rows.Next() {
		var status dto.Status

		err = rows.Scan(
			&status.ID,
			&status.BoardID,
			&status.Name,
			&status.Color,
			&status.CreatedAt,
			&status.UpdatedAt,
			&status.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan status: %w", err)
		}

		statuses = append(statuses, status)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while iterating statuses: %w", err)
	}

	return statuses, nil
}
