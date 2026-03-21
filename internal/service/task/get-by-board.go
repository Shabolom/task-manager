package task_service

import (
	"context"
	"task_manager/internal/dto"
)

func (s *Service) ListTasksByBoard(ctx context.Context, boardID int64) ([]dto.Task, error) {
	return s.taskRepo.ListTasksByBoard(ctx, boardID)
}
