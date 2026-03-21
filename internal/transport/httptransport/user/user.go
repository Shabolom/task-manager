package user_handler

import (
	"context"
	"task_manager/internal/dto"
	"time"
)

type UserService interface {
	CreateUser(ctx context.Context, user *dto.User) error
	DeleteUser(ctx context.Context, userID int64) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*dto.User, error)
	GetUserByID(ctx context.Context, userID int64) (*dto.User, error)
	ListUsers(ctx context.Context) ([]dto.User, error)
	UpdateUser(ctx context.Context, userID int64, request *dto.UpdateUserRequest) (*time.Time, error)
	UpdateUserPassword(ctx context.Context, userID int64, passwordHash string) (bool, error)
}

type Handler struct {
	userService UserService
	jwtSecret   string
}

func New(userRepo UserService, jwtSecret string) *Handler {
	return &Handler{userService: userRepo, jwtSecret: jwtSecret}
}
