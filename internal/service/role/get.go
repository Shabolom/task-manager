package role

import (
	"context"
	"task_manager/internal/dto"
)

func (s *Service) Get(ctx context.Context, roleID int64) (*dto.Role, error) {
	id, err := s.roleRepo.GetRoleByID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	return id, nil
}
