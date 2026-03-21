package auth_handler

import (
	"net/http"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h Handler) PostAuthLogin(ctx echo.Context) error {
	request := new(dto.LoginRequest)

	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	tokens, err := h.authService.Login(ctx.Request().Context(), request.Email, request.Password)
	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, "invalid credentials", err)
	}

	return render.JSON(ctx, http.StatusOK, tokens)
}
