package task_service

import (
	"context"
	"task_manager/internal/dto"
)

func (s *Service) ListTasksByStatus(ctx context.Context, statusID int64, pagination dto.Pagination) ([]dto.Task, error) {
	return s.taskRepo.ListTasksByStatus(ctx, statusID, pagination)
}
