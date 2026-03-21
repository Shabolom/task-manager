package user_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteUsersUserId(ctx echo.Context, userId api.UserId) error {
	id := int64(userId)
	deleted, err := h.userService.DeleteUser(ctx.Request().Context(), id)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}
	if !deleted {
		return render.BadRequest(ctx, render.CodeBadRequest, "User not found", nil)
	}

	return render.JSON(ctx, http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}
