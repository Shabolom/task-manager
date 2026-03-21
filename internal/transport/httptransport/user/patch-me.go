package user_handler

import (
	"net/http"
	"task_manager/internal/dto"
	"task_manager/internal/render"
	"task_manager/pkg/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PatchUsersMe(ctx echo.Context) error {
	userID, err := utils.ExtractUserIDFromToken(
		ctx.Request().Header.Get("Authorization"),
		h.jwtSecret,
	)

	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	request := new(dto.UpdateUserRequest)

	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	updatedAt, err := h.userService.UpdateUser(ctx.Request().Context(), userID, request)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusOK, map[string]interface{}{
		"user_id":    userID,
		"updated_at": updatedAt,
	})
}
