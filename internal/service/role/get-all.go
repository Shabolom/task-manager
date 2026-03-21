package role

import (
	"context"
	"task_manager/internal/dto"
)

func (s *Service) GetAll(ctx context.Context) ([]dto.Role, error) {
	roles, err := s.roleRepo.ListRoles(ctx)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
