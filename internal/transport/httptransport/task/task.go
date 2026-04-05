package task_handler

import (
	"context"
	"task_manager/internal/dto"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *dto.Task) error
	DeleteTask(ctx context.Context, taskID int64) (bool, error)
	GetTaskByID(ctx context.Context, taskID int64) (*dto.Task, error)
	ListTasksByBoard(ctx context.Context, boardID int64, pagination dto.Pagination) ([]dto.Task, error)
	ListTasksByStatus(ctx context.Context, statusID int64, pagination dto.Pagination) ([]dto.Task, error)
	UpdateTask(ctx context.Context, updateRequest dto.TaskUpdateRequest) error
}

type Handler struct {
	taskService TaskService
	jwtSecret   string
}

func New(taskService TaskService, jwtSecret string) *Handler {
	return &Handler{taskService: taskService, jwtSecret: jwtSecret}
}
