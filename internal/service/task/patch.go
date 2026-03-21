package task_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"
)

func (s *Service) UpdateTask(ctx context.Context, updateRequest dto.TaskUpdateRequest) error {
	if updateRequest.TaskID == 0 {
		return errors.New("taskID cannot be zero")
	}

	_, err := s.taskRepo.UpdateTask(ctx, updateRequest)
	
	return err
}
