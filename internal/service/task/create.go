package task_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"
	"time"
)

func (s *Service) CreateTask(ctx context.Context, task *dto.Task) error {
	if task.Title == "" {
		return errors.New("title cannot be empty")
	}
	if task.BoardID == 0 {
		return errors.New("boardID cannot be zero")
	}
	if task.StatusID == 0 {
		return errors.New("statusID cannot be zero")
	}

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	return s.taskRepo.CreateTask(ctx, task)
}
