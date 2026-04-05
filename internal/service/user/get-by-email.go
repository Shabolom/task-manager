package user_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*dto.User, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		s.logger.Warn("Failed to GetUserByEmail", zap.Error(err))
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
