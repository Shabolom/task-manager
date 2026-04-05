package task_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) UpdateTask(ctx context.Context, updateRequest dto.TaskUpdateRequest) error {
	if updateRequest.TaskID == 0 {
		return errors.New("taskID cannot be zero")
	}

	_, err := s.taskRepo.UpdateTask(ctx, updateRequest)
	if err != nil {
		s.logger.Warn("Failed to UpdateTask", zap.Error(err))
		return err
	}

	return nil
}
