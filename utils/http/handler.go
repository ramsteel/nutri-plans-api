package http

import (
	"net/http"
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

func HandleSearchResponse(
	c echo.Context,
	message string,
	data any,
	metadata *dto.MetadataResponse,
) error {
	return c.JSON(http.StatusOK, &dto.SearchResponse{
		BaseResponse: dto.BaseResponse{
			Status:  statusconst.StatusSuccess,
			Message: message,
			Data:    data,
		},
		Metadata: metadata,
	})
}

func HandlePaginationResponse(
	c echo.Context,
	message string,
	data any,
	pagination *dto.PaginationMetadata,
	link *dto.Link,
) error {
	return c.JSON(http.StatusOK, &dto.PaginationResponse{
		BaseResponse: dto.BaseResponse{
			Status:  statusconst.StatusSuccess,
			Message: message,
			Data:    data,
		},
		Pagination: pagination,
		Link:       link,
	})
}
