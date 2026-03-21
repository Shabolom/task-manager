package statuses

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) CreateStatus(ctx context.Context, status *dto.Status) error {
	query := `
		INSERT INTO statuses (
			board_id,
			name,
			color
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			created_at,
			updated_at,
			deleted_at
	`

	err := s.conn.QueryRow(ctx, query,
		status.BoardID,
		status.Name,
		status.Color,
	).Scan(
		&status.ID,
		&status.CreatedAt,
		&status.UpdatedAt,
		&status.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create status: %w", err)
	}

	return nil
}
