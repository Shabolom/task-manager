package status_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PatchStatusesStatusId(ctx echo.Context, statusId api.StatusId) error {
	request := new(dto.Status)
	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	updatedAt, err := h.statusService.UpdateStatus(ctx.Request().Context(), int64(statusId), request.Name, request.Color)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusOK, map[string]interface{}{
		"status_id":  int64(statusId),
		"updated_at": updatedAt,
	})
}
