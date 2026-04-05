package board

import (
	"context"
	"task_manager/internal/dto"
)

func (s *Service) ListBoardsByOwner(ctx context.Context, ownerID int64, pagination dto.Pagination) ([]dto.Board, error) {
	return s.boardRepo.ListBoardsByOwner(ctx, ownerID, pagination)
}
