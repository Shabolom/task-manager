package board

import (
	"context"
	"errors"
	"task_manager/internal/dto"

	"go.uber.org/zap"
)

func (s *Service) GetBoardByID(ctx context.Context, boardID int64) (*dto.Board, error) {
	board, err := s.boardRepo.GetBoardByID(ctx, boardID)
	if err != nil {
		s.logger.Warn("Failed to get board", zap.Error(err))
		return nil, err
	}
	if board == nil {
		return nil, errors.New("board not found")
	}
	return board, nil
}
