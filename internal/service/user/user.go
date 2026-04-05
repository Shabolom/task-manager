package user_service

import (
	"context"
	"task_manager/internal/dto"
	"time"

	"go.uber.org/zap"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *dto.User) error
	DeleteUser(ctx context.Context, userID int64) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*dto.User, error)
	GetUserByID(ctx context.Context, userID int64) (*dto.User, error)
	ListUsers(ctx context.Context, pagination dto.Pagination) ([]dto.User, error)
	UpdateUser(ctx context.Context, userID int64, request *dto.UpdateUserRequest) (*time.Time, error)
	UpdateUserPassword(ctx context.Context, userID int64, passwordHash []byte) (int64, error)
}

type Service struct {
	userRepo UserRepo
	logger   *zap.Logger
}

func New(userRepo UserRepo, logger *zap.Logger) *Service {
	return &Service{userRepo: userRepo, logger: logger}
}
