package status_service

import (
	"context"
	"task_manager/internal/dto"
)

func (s *Service) ListStatusesByBoard(ctx context.Context, boardID int64) ([]dto.Status, error) {
	return s.statusRepo.ListStatusesByBoard(ctx, boardID)
}
