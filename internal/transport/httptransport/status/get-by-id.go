package status_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetStatusesStatusId(ctx echo.Context, statusId api.StatusId) error {
	status, err := h.statusService.GetStatusByID(ctx.Request().Context(), int64(statusId))
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}
	return render.JSON(ctx, http.StatusOK, status)
}
