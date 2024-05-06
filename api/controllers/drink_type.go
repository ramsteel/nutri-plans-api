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

type drinkTypeController struct {
	drinkTypeUsecase usecases.DrinkTypeUsecase
}

func NewDrinkTypeController(drinkTypeUsecase usecases.DrinkTypeUsecase) *drinkTypeController {
	return &drinkTypeController{
		drinkTypeUsecase: drinkTypeUsecase,
	}
}

func (f *drinkTypeController) GetDrinkTypes(c echo.Context) error {
	res, err := f.drinkTypeUsecase.GetDrinkTypes(c)
	if err != nil {
		var (
			code int    = http.StatusInternalServerError
			msg  string = msgconst.MsgGetDrinkTypesFailed
		)

		if errors.Is(err, context.Canceled) {
			code = httpconst.StatusClientCancelledRequest
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgRetrieveDrinkTypesSuccess,
		res,
	)
}
