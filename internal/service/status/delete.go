package status_service

import (
	"context"

	"go.uber.org/zap"
)

func (s *Service) DeleteStatus(ctx context.Context, statusID int64) (bool, error) {
	rowsAffected, err := s.statusRepo.DeleteStatus(ctx, statusID)
	if err != nil {
		s.logger.Warn("Failed to deleteStatus", zap.Error(err))
		return false, err
	}
	return rowsAffected > 0, nil
}
