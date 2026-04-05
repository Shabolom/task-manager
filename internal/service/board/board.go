package board

import (
	"context"
	"task_manager/internal/dto"
	"time"

	"go.uber.org/zap"
)

type BoardRepo interface {
	CreateBoard(ctx context.Context, board *dto.Board) error
	DeleteBoard(ctx context.Context, boardID int64) (int64, error)
	GetBoardByID(ctx context.Context, boardID int64) (*dto.Board, error)
	ListBoardsByOwner(ctx context.Context, ownerID int64) ([]dto.Board, error)
	UpdateBoard(ctx context.Context, boardID int64, title, description string) (*time.Time, error)
}

type Service struct {
	boardRepo BoardRepo
	logger    *zap.Logger
}

func New(boardRepo BoardRepo, logger *zap.Logger) *Service {
	return &Service{boardRepo: boardRepo, logger: logger}
}
