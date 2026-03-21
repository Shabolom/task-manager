package status_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteStatusesStatusId(ctx echo.Context, statusId api.StatusId) error {
	deleted, err := h.statusService.DeleteStatus(ctx.Request().Context(), int64(statusId))
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}
	if !deleted {
		return render.BadRequest(ctx, render.CodeBadRequest, "Status not found", nil)
	}

	return render.JSON(ctx, http.StatusOK, map[string]string{
		"message": "Status deleted successfully",
	})
}
