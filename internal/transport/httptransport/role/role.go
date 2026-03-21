package role

import (
	"context"
	"task_manager/internal/dto"
)

type RoleService interface {
	Get(ctx context.Context, roleID int64) (*dto.Role, error)
	GetAll(ctx context.Context) ([]dto.Role, error)
	Create(ctx context.Context, role *dto.Role) error
	DeleteRole(ctx context.Context, roleID int64) (int64, error)
	Patch(ctx context.Context, patchReq dto.UpdateRoleRequest) (dto.Role, error)
}

type Handler struct {
	roleService RoleService
}

func New(roleRepo RoleService) *Handler {
	return &Handler{roleService: roleRepo}
}
