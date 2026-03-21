package board_handler

import (
	"context"
	"task_manager/internal/dto"
	"time"
)

type BoardService interface {
	CreateBoard(ctx context.Context, board *dto.Board) error
	DeleteBoard(ctx context.Context, boardID int64) (bool, error)
	GetBoardByID(ctx context.Context, boardID int64) (*dto.Board, error)
	ListBoardsByOwner(ctx context.Context, ownerID int64) ([]dto.Board, error)
	UpdateBoard(ctx context.Context, boardID int64, title, description string) (*time.Time, error)
}

type Handler struct {
	boardService BoardService
	jwtSecret    string
}

func New(boardRepo BoardService, jwtSecret string) *Handler {
	return &Handler{boardService: boardRepo, jwtSecret: jwtSecret}
}
