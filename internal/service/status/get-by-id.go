package status_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) GetStatusByID(ctx context.Context, statusID int64) (*dto.Status, error) {
	status, err := s.statusRepo.GetStatusByID(ctx, statusID)
	if err != nil {
		s.logger.Warn("Failed to GetStatusByID", zap.Error(err))
		return nil, err
	}
	if status == nil {
		return nil, errors.New("status not found")
	}
	return status, nil
}
