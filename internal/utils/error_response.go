package utils

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	error string `json:"error"`
}

func SendErrorResponse(ctx echo.Context, msg string, statusCode int) error {
	response := ErrorResponse{
		error: msg,
	}
	return ctx.JSON(statusCode, response)
}
