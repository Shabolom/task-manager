package user_service

import (
	"context"
	"task_manager/internal/dto"
)

func (s *Service) ListUsers(ctx context.Context, pagination dto.Pagination) ([]dto.User, error) {
	return s.userRepo.ListUsers(ctx, pagination)
}
