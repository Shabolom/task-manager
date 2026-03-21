package user_handler

import (
	"net/http"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PostUsers(ctx echo.Context) error {
	request := new(dto.User)
	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	if ctx.Echo().Validator != nil {
		if err := ctx.Validate(request); err != nil {
			return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
		}
	}

	if err := h.userService.CreateUser(ctx.Request().Context(), request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusCreated, request)
}
