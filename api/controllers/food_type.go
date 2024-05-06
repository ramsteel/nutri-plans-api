package controllers

import (
	"context"
	"errors"
	"net/http"
	httpconst "nutri-plans-api/constants/http"
	msgconst "nutri-plans-api/constants/message"
	"nutri-plans-api/usecases"
	httputil "nutri-plans-api/utils/http"

	"github.com/labstack/echo/v4"
)

type foodTypeController struct {
	foodTypeUsecase usecases.FoodTypeUsecase
}

func NewFoodTypeController(foodTypeUsecase usecases.FoodTypeUsecase) *foodTypeController {
	return &foodTypeController{
		foodTypeUsecase: foodTypeUsecase,
	}
}

func (f *foodTypeController) GetFoodTypes(c echo.Context) error {
	res, err := f.foodTypeUsecase.GetFoodTypes(c)
	if err != nil {
		var (
			code int    = http.StatusInternalServerError
			msg  string = msgconst.MsgGetFoodTypesFailed
		)

		if errors.Is(err, context.Canceled) {
			code = httpconst.StatusClientCancelledRequest
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgRetrieveFoodTypesSuccess,
		res,
	)
}
