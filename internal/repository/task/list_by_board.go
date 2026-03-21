package task

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) ListTasksByBoard(ctx context.Context, boardID int64) ([]dto.Task, error) {
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
		WHERE board_id = $1
		  AND deleted_at IS NULL
		ORDER BY position ASC, id ASC
	`

	rows, err := s.conn.Query(ctx, query, boardID)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks by board %d: %w", boardID, err)
	}
	defer rows.Close()

	tasks := make([]dto.Task, 0)

	for rows.Next() {
		var task dto.Task

		err = rows.Scan(
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
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while iterating tasks: %w", err)
	}

	return tasks, nil
}
