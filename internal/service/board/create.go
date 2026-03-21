package board

import (
	"context"
	"errors"
	"task_manager/internal/dto"
	"time"
)

func (s *Service) CreateBoard(ctx context.Context, board *dto.Board) error {
	if board.Title == "" {
		return errors.New("title cannot be empty")
	}
	if board.OwnerID == 0 {
		return errors.New("ownerID cannot be zero")
	}

	board.CreatedAt = time.Now()
	board.UpdatedAt = time.Now()

	return s.boardRepo.CreateBoard(ctx, board)
}
