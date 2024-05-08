package controllers

import (
	"context"
	"errors"
	"net/http"
	httpconst "nutri-plans-api/constants/http"
	msgconst "nutri-plans-api/constants/message"
	"nutri-plans-api/dto"
	"nutri-plans-api/usecases"
	errutil "nutri-plans-api/utils/error"
	httputil "nutri-plans-api/utils/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type nutritionController struct {
	nutritionUsecase usecases.NutritionUsecase
}

func NewNutritionController(nutritionUsecase usecases.NutritionUsecase) *nutritionController {
	return &nutritionController{
		nutritionUsecase: nutritionUsecase,
	}
}

func (n *nutritionController) SearchItem(c echo.Context) error {
	item := strings.TrimSpace(c.QueryParam("item"))
	offset := c.QueryParam("offset")
	limit := c.QueryParam("limit")

	if !n.isValidItem(item) {
		return httputil.HandleErrorResponse(c, http.StatusBadRequest, msgconst.MsgQueryMinimum)
	}

	intOffset, intLimit, ok := n.parseQuery(offset, limit)
	if !ok {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	res, meta, err := n.nutritionUsecase.SearchItem(c, item, intLimit, intOffset)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, errutil.ErrExternalService):
			code = http.StatusBadGateway
			msg = msgconst.MsgExternalServiceError
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgSearchItemFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgSearchItemFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSearchResponse(
		c,
		http.StatusOK,
		msgconst.MsgRetrieveItemSuccess,
		res,
		meta,
	)
}

func (n *nutritionController) GetItemNutrition(c echo.Context) error {
	item := strings.TrimSpace(c.Param("item-name"))

	r := &dto.ItemNutritionRequest{
		Query: item,
	}
	res, err := n.nutritionUsecase.GetItemNutrition(c, r)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, errutil.ErrItemNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgItemNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgGetItemNutritionFailed
		case errors.Is(err, errutil.ErrExternalService):
			code = http.StatusBadGateway
			msg = msgconst.MsgExternalServiceError
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgGetItemNutritionFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgGetItemNutritionSuccess,
		res,
	)
}

func (n *nutritionController) parseQuery(offset, limit string) (int, int, bool) {
	if strings.TrimSpace(limit) == "" {
		limit = "10"
	}

	if strings.TrimSpace(offset) == "" {
		offset = "0"
	}

	intOffset, err := strconv.Atoi(offset)
	if err != nil {
		return 0, 0, false
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		return 0, 0, false
	}

	return intOffset, intLimit, true
}

func (n *nutritionController) isValidItem(item string) bool {
	return len(item) >= 3
}
