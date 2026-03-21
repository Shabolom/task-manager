package task_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteTasksTaskId(ctx echo.Context, taskId api.TaskId) error {
	deleted, err := h.taskService.DeleteTask(ctx.Request().Context(), int64(taskId))
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	if !deleted {
		return render.BadRequest(ctx, render.CodeBadRequest, "Task not found", nil)
	}

	return render.JSON(ctx, http.StatusOK, map[string]string{
		"message": "Task deleted successfully",
	})
}
