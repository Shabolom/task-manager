package role

import (
	"context"
	"task_manager/internal/dto"

	"go.uber.org/zap"
)

type RoleRepo interface {
	GetRoleByID(ctx context.Context, roleID int64) (*dto.Role, error)
	ListRoles(ctx context.Context, pagination dto.Pagination) ([]dto.Role, error)
	CreateRole(ctx context.Context, role *dto.Role) error
	DeleteRole(ctx context.Context, roleID int64) (int64, error)
	UpdateRole(ctx context.Context, request dto.UpdateRoleRequest) (dto.Role, error)
}

type Service struct {
	roleRepo RoleRepo
	logger   *zap.Logger
}

func New(roleRepo RoleRepo, logger *zap.Logger) *Service {
	return &Service{
		roleRepo: roleRepo,
		logger:   logger}
}
