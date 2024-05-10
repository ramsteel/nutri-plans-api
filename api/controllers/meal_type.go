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

type mealTypeController struct {
	mealTypeUsecase usecases.MealTypeUsecase
}

func NewMealTypeController(mealTypeUsecase usecases.MealTypeUsecase) *mealTypeController {
	return &mealTypeController{
		mealTypeUsecase: mealTypeUsecase,
	}
}

func (f *mealTypeController) GetMealTypes(c echo.Context) error {
	res, err := f.mealTypeUsecase.GetMealTypes(c)
	if err != nil {
		var (
			code int    = http.StatusInternalServerError
			msg  string = msgconst.MsgGetMealTypesFailed
		)

		if errors.Is(err, context.Canceled) {
			code = httpconst.StatusClientCancelledRequest
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgRetrieveMealTypesSuccess,
		res,
	)
}
