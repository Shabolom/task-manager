package user_service

import (
	"context"

	"go.uber.org/zap"
)

func (s *Service) DeleteUser(ctx context.Context, userID int64) (bool, error) {
	rowsAffected, err := s.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		s.logger.Warn("Failed to DeleteUser", zap.Error(err))
		return false, err
	}
	return rowsAffected > 0, nil
}
