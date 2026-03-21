package board

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) GetBoardByID(ctx context.Context, boardID int64) (*dto.Board, error) {
	query := `
		SELECT
			id,
			title,
			description,
			owner_id,
			created_at,
			updated_at,
			deleted_at
		FROM boards
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	var board dto.Board

	err := s.conn.QueryRow(ctx, query, boardID).Scan(
		&board.ID,
		&board.Title,
		&board.Description,
		&board.OwnerID,
		&board.CreatedAt,
		&board.UpdatedAt,
		&board.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get board by id %d: %w", boardID, err)
	}

	return &board, nil
}
