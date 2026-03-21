package board

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) CreateBoard(ctx context.Context, board *dto.Board) error {
	query := `
		INSERT INTO boards (
			title,
			description,
			owner_id
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			created_at,
			updated_at,
			deleted_at
	`

	err := s.conn.QueryRow(ctx, query,
		board.Title,
		board.Description,
		board.OwnerID,
	).Scan(
		&board.ID,
		&board.CreatedAt,
		&board.UpdatedAt,
		&board.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create board: %w", err)
	}

	return nil
}
