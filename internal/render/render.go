package render

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	CodeBadRequest          = "BAD_REQUEST"
	CodeValidationError     = "VALIDATION_ERROR"
	CodeUnauthorized        = "UNAUTHORIZED"
	CodeForbidden           = "FORBIDDEN"
	CodeNotFound            = "NOT_FOUND"
	CodeConflict            = "CONFLICT"
	CodeUnprocessableEntity = "UNPROCESSABLE_ENTITY"
	CodeTooManyRequests     = "TOO_MANY_REQUESTS"
	CodeInternalError       = "INTERNAL_ERROR"
	CodeServiceUnavailable  = "SERVICE_UNAVAILABLE"
)

const (
	MsgBadRequest          = "Bad request"
	MsgValidationError     = "Validation failed"
	MsgUnauthorized        = "Unauthorized"
	MsgForbidden           = "Forbidden"
	MsgNotFound            = "Resource not found"
	MsgConflict            = "Conflict"
	MsgUnprocessableEntity = "Unprocessable entity"
	MsgTooManyRequests     = "Too many requests"
	MsgInternalError       = "Internal server error"
	MsgServiceUnavailable  = "Service unavailable"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func JSON(c echo.Context, status int, body any) error {
	return c.JSON(status, body)
}

func Error(c echo.Context, status int, code, message string, err error, exposeErr bool) error {
	resp := ErrorResponse{
		Code:    code,
		Message: message,
	}

	if err != nil && exposeErr {
		resp.Error = err.Error()
	}

	return c.JSON(status, resp)
}

func BadRequest(c echo.Context, code, message string, err error) error {
	if code == "" {
		code = CodeBadRequest
	}
	if message == "" {
		message = MsgBadRequest
	}

	return Error(c, http.StatusBadRequest, code, message, err, true)
}

func ValidationError(c echo.Context, err error) error {
	return Error(
		c,
		http.StatusBadRequest,
		CodeValidationError,
		MsgValidationError,
		err,
		true,
	)
}

func Unauthorized(c echo.Context, err error) error {
	return Error(
		c,
		http.StatusUnauthorized,
		CodeUnauthorized,
		MsgUnauthorized,
		err,
		false,
	)
}

func Forbidden(c echo.Context, err error) error {
	return Error(
		c,
		http.StatusForbidden,
		CodeForbidden,
		MsgForbidden,
		err,
		false,
	)
}

func NotFound(c echo.Context, code, message string, err error) error {
	if code == "" {
		code = CodeNotFound
	}
	if message == "" {
		message = MsgNotFound
	}

	return Error(c, http.StatusNotFound, code, message, err, false)
}

func Conflict(c echo.Context, code, message string, err error) error {
	if code == "" {
		code = CodeConflict
	}
	if message == "" {
		message = MsgConflict
	}

	return Error(c, http.StatusConflict, code, message, err, true)
}

func UnprocessableEntity(c echo.Context, code, message string, err error) error {
	if code == "" {
		code = CodeUnprocessableEntity
	}
	if message == "" {
		message = MsgUnprocessableEntity
	}

	return Error(c, http.StatusUnprocessableEntity, code, message, err, true)
}

func TooManyRequests(c echo.Context, err error) error {
	return Error(
		c,
		http.StatusTooManyRequests,
		CodeTooManyRequests,
		MsgTooManyRequests,
		err,
		false,
	)
}

func Internal(c echo.Context, err error) error {
	return Error(
		c,
		http.StatusInternalServerError,
		CodeInternalError,
		MsgInternalError,
		err,
		false,
	)
}

func ServiceUnavailable(c echo.Context, err error) error {
	return Error(
		c,
		http.StatusServiceUnavailable,
		CodeServiceUnavailable,
		MsgServiceUnavailable,
		err,
		false,
	)
}
