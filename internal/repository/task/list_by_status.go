package task

import (
	"context"
	"fmt"
	"task_manager/internal/dto"
)

func (s *Storage) ListTasksByStatus(ctx context.Context, statusID int64, pagination dto.Pagination) ([]dto.Task, error) {
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
		WHERE status_id = $1
		  AND deleted_at IS NULL
		ORDER BY position ASC, id ASC
		LIMIT $2 OFFSET $3
	`

	rows, err := s.conn.Query(ctx, query, statusID, pagination.Limit, pagination.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks by status %d: %w", statusID, err)
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
