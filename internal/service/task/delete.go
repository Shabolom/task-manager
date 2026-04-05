package task_service

import (
	"context"

	"go.uber.org/zap"
)

func (s *Service) DeleteTask(ctx context.Context, taskID int64) (bool, error) {
	rowsAffected, err := s.taskRepo.DeleteTask(ctx, taskID)
	if err != nil {
		s.logger.Warn("Failed to DeleteTask", zap.Error(err))
		return false, err
	}
	return rowsAffected > 0, nil
}
