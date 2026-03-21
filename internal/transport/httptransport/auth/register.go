package auth_handler

import (
	"net/http"
	"task_manager/internal/dto"
	"task_manager/internal/render"

	"github.com/labstack/echo/v4"
)

func (h Handler) PostAuthRegister(ctx echo.Context) error {
	request := new(dto.RegisterRequest)

	if err := ctx.Bind(request); err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	if ctx.Echo().Validator != nil {
		if err := ctx.Validate(request); err != nil {
			return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
		}
	}

	err := h.authService.Register(
		ctx.Request().Context(),
		request.Email,
		request.Password,
		request.FullName,
	)

	if err != nil {
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusCreated, map[string]string{
		"message": "user created",
	})
}
