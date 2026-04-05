package task_service

import (
	"context"
	"task_manager/internal/dto"
	"time"

	"go.uber.org/zap"
)

type TaskRepo interface {
	CreateTask(ctx context.Context, task *dto.Task) error
	DeleteTask(ctx context.Context, taskID int64) (int64, error)
	GetTaskByID(ctx context.Context, taskID int64) (*dto.Task, error)
	ListTasksByBoard(ctx context.Context, boardID int64, pagination dto.Pagination) ([]dto.Task, error)
	ListTasksByStatus(ctx context.Context, statusID int64, pagination dto.Pagination) ([]dto.Task, error)
	UpdateTask(ctx context.Context, updateRequest dto.TaskUpdateRequest) (*time.Time, error)
}

type Service struct {
	taskRepo TaskRepo
	logger   *zap.Logger
}

func New(taskRepo TaskRepo, logger *zap.Logger) *Service {
	return &Service{taskRepo: taskRepo, logger: logger}
}
