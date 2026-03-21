package user_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"
	"time"
)

func (s *Service) CreateUser(ctx context.Context, user *dto.User) error {
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}
	if user.FullName == "" {
		return errors.New("fullName cannot be empty")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return s.userRepo.CreateUser(ctx, user)
}
