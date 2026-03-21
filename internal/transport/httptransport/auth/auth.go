package auth_handler

import (
	"context"
	"task_manager/internal/dto"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (*dto.TokenPair, error)
	Register(ctx context.Context, email, password, fullName string) error
	Refresh(ctx context.Context, refreshToken string) (*dto.TokenPair, error)
}

type Handler struct {
	authService AuthService
}

func New(authService AuthService) *Handler {
	return &Handler{authService: authService}
}
