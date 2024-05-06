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

type dietaryPreferenceTypeController struct {
	dietaryPreferenceTypeUsecase usecases.DietaryPreferenceTypeUsecase
}

func NewDietaryPreferenceTypeController(dietaryPreferenceTypeUsecase usecases.DietaryPreferenceTypeUsecase) *dietaryPreferenceTypeController {
	return &dietaryPreferenceTypeController{
		dietaryPreferenceTypeUsecase: dietaryPreferenceTypeUsecase,
	}
}

func (f *dietaryPreferenceTypeController) GetDietaryPreferenceTypes(c echo.Context) error {
	res, err := f.dietaryPreferenceTypeUsecase.GetDietaryPreferenceTypes(c)
	if err != nil {
		var (
			code int    = http.StatusInternalServerError
			msg  string = msgconst.MsgGetDietaryPreferenceTypesFailed
		)

		if errors.Is(err, context.Canceled) {
			code = httpconst.StatusClientCancelledRequest
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgRetrieveDietaryPreferenceTypesSuccess,
		res,
	)
}
