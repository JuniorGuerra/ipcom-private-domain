package utils

import (
	"github.com/labstack/echo/v4"
)

type ErrorSMS struct {
	StatusCode int
	Message    string
	Data       map[string]interface{}
}

func ErrorResponse(errorsms ErrorSMS, c echo.Context) error {
	return c.JSON(errorsms.StatusCode, errorsms)
}
