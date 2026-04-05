package task_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) GetTaskByID(ctx context.Context, taskID int64) (*dto.Task, error) {
	task, err := s.taskRepo.GetTaskByID(ctx, taskID)
	if err != nil {
		s.logger.Warn("Failed to GetTaskByID", zap.Error(err))
		return nil, err
	}
	if task == nil {
		return nil, errors.New("task not found")
	}
	return task, nil
}
