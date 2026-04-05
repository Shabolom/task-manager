package board

import (
	"context"

	"go.uber.org/zap"
)

func (s *Service) DeleteBoard(ctx context.Context, boardID int64) (bool, error) {
	rowsAffected, err := s.boardRepo.DeleteBoard(ctx, boardID)
	if err != nil {
		s.logger.Warn("Failed to delete board", zap.Error(err))
		return false, err
	}
	return rowsAffected > 0, nil
}
