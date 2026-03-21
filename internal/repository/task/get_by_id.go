package task

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) GetTaskByID(ctx context.Context, taskID int64) (*dto.Task, error) {
	query := `
		SELECT
			id,
			board_id,
			status_id,
			title,
			description,
			priority,
			due_date,
			created_at,
			updated_at,
			deleted_at
		FROM tasks
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	var task dto.Task

	err := s.conn.QueryRow(ctx, query, taskID).Scan(
		&task.ID,
		&task.BoardID,
		&task.StatusID,
		&task.Title,
		&task.Description,
		&task.Priority,
		&task.DueDate,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get task by id %d: %w", taskID, err)
	}

	return &task, nil
}
