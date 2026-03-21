package user_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PatchUsersUserIdPassword(ctx echo.Context, userId api.UserId) error {
	id := int64(userId)
	request := new(dto.UpdateUserPasswordRequest)
	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	success, err := h.userService.UpdateUserPassword(ctx.Request().Context(), id, request.Password)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}
	if !success {
		return render.BadRequest(ctx, render.CodeBadRequest, "User not found or password not updated", nil)
	}

	return render.JSON(ctx, http.StatusOK, map[string]string{
		"message": "Password updated successfully",
	})
}
