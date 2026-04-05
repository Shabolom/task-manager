package user_service

import (
	"context"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) UpdateUserPassword(ctx context.Context, userID int64, passwordHash string) (bool, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordHash), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Warn("Failed to GenerateFromPassword", zap.Error(err))
		return false, err
	}

	rowsAffected, err := s.userRepo.UpdateUserPassword(ctx, userID, hash)
	if err != nil {
		s.logger.Warn("Failed to UpdateUserPassword", zap.Error(err))
		return false, err
	}
	return rowsAffected > 0, nil
}
