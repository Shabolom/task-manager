package user_handler

import (
	"net/http"
	"task_manager/internal/render"
	"task_manager/pkg/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUsersMe(ctx echo.Context) error {
	userID, err := utils.ExtractUserIDFromToken(
		ctx.Request().Header.Get("Authorization"),
		h.jwtSecret,
	)

	user, err := h.userService.GetUserByID(ctx.Request().Context(), userID)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}
	return render.JSON(ctx, http.StatusOK, user)
}
