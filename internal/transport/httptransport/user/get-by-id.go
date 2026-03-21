package user_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUsersUserId(ctx echo.Context, userId api.UserId) error {
	id := int64(userId)
	user, err := h.userService.GetUserByID(ctx.Request().Context(), id)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}
	return render.JSON(ctx, http.StatusOK, user)
}
