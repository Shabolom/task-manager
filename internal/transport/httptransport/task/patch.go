package task_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PatchTasksTaskId(ctx echo.Context, taskId api.TaskId) error {
	request := new(dto.TaskUpdateRequest)

	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	request.TaskID = int64(taskId)

	if ctx.Echo().Validator != nil {
		if err := ctx.Validate(request); err != nil {
			return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
		}
	}

	err := h.taskService.UpdateTask(ctx.Request().Context(), *request)
	if err != nil {
		return render.Internal(ctx, err)
	}

	task, err := h.taskService.GetTaskByID(ctx.Request().Context(), request.TaskID)
	if err != nil {
		return render.Internal(ctx, err)
	}

	return render.JSON(ctx, http.StatusOK, task)
}
