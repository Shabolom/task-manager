package user_handler

import (
	"net/http"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUsers(ctx echo.Context) error {
	users, err := h.userService.ListUsers(ctx.Request().Context())
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}
	return render.JSON(ctx, http.StatusOK, users)
}
