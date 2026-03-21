package auth_handler

import (
	"net/http"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h Handler) PostAuthRefresh(ctx echo.Context) error {
	request := new(dto.RefreshRequest)

	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	tokens, err := h.authService.Refresh(ctx.Request().Context(), request.RefreshToken)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, "invalid refresh token", err)
	}

	return render.JSON(ctx, http.StatusOK, tokens)
}
