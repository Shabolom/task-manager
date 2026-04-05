package di

import (
	board_repo "task_manager/internal/repository/board"
	role_repo "task_manager/internal/repository/role"
	status_repo "task_manager/internal/repository/statuses"
	task_repo "task_manager/internal/repository/task"
	user_repo "task_manager/internal/repository/user"
)

func (d *DI) GetRoleRepo() *role_repo.Storage {
	return role_repo.New(d.GetPgDatabase())
}

func (d *DI) GetBoardRepo() *board_repo.Storage {
	return board_repo.New(d.GetPgDatabase())
}

func (d *DI) GetUserRepo() *user_repo.Storage {
	return user_repo.New(d.GetPgDatabase())
}

func (d *DI) GetStatusRepo() *status_repo.Storage {
	return status_repo.New(d.GetPgDatabase())
}

func (d *DI) GetTaskRepo() *task_repo.Storage {
	return task_repo.New(d.GetPgDatabase())
}
