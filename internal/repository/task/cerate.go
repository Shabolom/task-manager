package task

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) CreateTask(ctx context.Context, task *dto.Task) error {
	query := `
		INSERT INTO tasks (
			board_id,
			status_id,
			title,
			description,
			priority,
			due_date,
		    creator_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			id,
			created_at,
			updated_at,
			deleted_at
	`

	err := s.conn.QueryRow(ctx, query,
		task.BoardID,
		task.StatusID,
		task.Title,
		task.Description,
		task.Priority,
		task.DueDate,
		task.CreatorId,
	).Scan(
		&task.ID,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}
