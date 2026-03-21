package user_handler

import (
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PatchUsersUserId(ctx echo.Context, userId api.UserId) error {
	id := int64(userId)
	request := new(dto.UpdateUserRequest)
	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	updatedAt, err := h.userService.UpdateUser(ctx.Request().Context(), id, request)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusOK, map[string]interface{}{
		"user_id":    id,
		"updated_at": updatedAt,
	})
}
