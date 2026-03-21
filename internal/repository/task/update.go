package task

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
	"time"
)

func (s *Storage) UpdateTask(ctx context.Context, updateRequest dto.TaskUpdateRequest) (*time.Time, error) {
	query := `
		UPDATE tasks
		SET
			board_id = COALESCE($1, board_id),
			status_id = COALESCE($2, status_id),
			title = COALESCE($3, title),
			description = COALESCE($4, description),
			position = COALESCE($5, position),
			due_date = COALESCE($6, due_date),
			updated_at = NOW()
		WHERE id = $7
		  AND deleted_at IS NULL
		RETURNING updated_at
	`

	var updatedAt time.Time

	err := s.conn.QueryRow(ctx, query,
		updateRequest.BoardID,
		updateRequest.StatusID,
		updateRequest.Title,
		updateRequest.Description,
		updateRequest.Position,
		updateRequest.DueDate,
		updateRequest.TaskID,
	).Scan(&updatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to update task %d: %w", updateRequest.TaskID, err)
	}

	return &updatedAt, nil
}
