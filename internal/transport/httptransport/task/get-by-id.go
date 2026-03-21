package task_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetTasksTaskId(ctx echo.Context, taskId api.TaskId) error {
	task, err := h.taskService.GetTaskByID(ctx.Request().Context(), int64(taskId))
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusOK, task)
}
