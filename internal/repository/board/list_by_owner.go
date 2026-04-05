package board

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) ListBoardsByOwner(ctx context.Context, ownerID int64, pagination dto.Pagination) ([]dto.Board, error) {
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
		WHERE owner_id = $1
		  AND deleted_at IS NULL
		ORDER BY id ASC 
		LIMIT $2 OFFSET $3
	`

	rows, err := s.conn.Query(ctx, query, ownerID, pagination.Limit, pagination.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list boards by owner %d: %w", ownerID, err)
	}
	defer rows.Close()

	boards := make([]dto.Board, 0)

	for rows.Next() {
		var board dto.Board

		err = rows.Scan(
			&board.ID,
			&board.Title,
			&board.Description,
			&board.OwnerID,
			&board.CreatedAt,
			&board.UpdatedAt,
			&board.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan board: %w", err)
		}

		boards = append(boards, board)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while iterating boards: %w", err)
	}

	return boards, nil
}
