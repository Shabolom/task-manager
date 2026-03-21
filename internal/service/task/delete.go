package task_service

import "context"

func (s *Service) DeleteTask(ctx context.Context, taskID int64) (bool, error) {
	rowsAffected, err := s.taskRepo.DeleteTask(ctx, taskID)
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}
