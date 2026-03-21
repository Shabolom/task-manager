package task_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"
)

func (s *Service) GetTaskByID(ctx context.Context, taskID int64) (*dto.Task, error) {
	task, err := s.taskRepo.GetTaskByID(ctx, taskID)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, errors.New("task not found")
	}
	return task, nil
}
