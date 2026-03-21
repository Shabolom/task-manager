package user_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"
)

func (s *Service) GetUserByID(ctx context.Context, userID int64) (*dto.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
