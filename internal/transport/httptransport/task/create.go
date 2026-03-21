package task_handler

import (
	"fmt"
	"net/http"
	api "task_manager/gen"
	"task_manager/internal/dto"
	"task_manager/internal/render"
	"task_manager/pkg/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PostBoardsBoardIdTasks(ctx echo.Context, boardId api.BoardId) error {
	request := new(dto.Task)

	userID, err := utils.ExtractUserIDFromToken(
		ctx.Request().Header.Get("Authorization"),
		h.jwtSecret,
	)
	if err != nil {
		return render.Unauthorized(ctx, err)
	}

	if err := ctx.Bind(request); err != nil {
		fmt.Println(err)
		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	request.BoardID = int64(boardId)
	request.CreatorId = userID
	
	if ctx.Echo().Validator != nil {
		if err := ctx.Validate(request); err != nil {
			fmt.Println(err)
			return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
		}
	}

	if err := h.taskService.CreateTask(ctx.Request().Context(), request); err != nil {
		fmt.Println(err)

		return render.BadRequest(ctx, render.CodeBadRequest, render.MsgBadRequest, err)
	}

	return render.JSON(ctx, http.StatusCreated, request)
}
