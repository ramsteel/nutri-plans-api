package http

import (
	statusconst "nutri-plans-api/constants/status"
	"nutri-plans-api/dto"

	"github.com/labstack/echo/v4"
)

func HandleErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, &dto.BaseResponse{
		Status:  statusconst.StatusFailed,
		Message: message,
	})
}

func HandleSuccessResponse(c echo.Context, code int, message string, data any) error {
	return c.JSON(code, &dto.BaseResponse{
		Status:  statusconst.StatusSuccess,
		Message: message,
		Data:    data,
	})
}
