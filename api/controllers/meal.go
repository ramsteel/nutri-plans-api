package controllers

import (
	"context"
	"errors"
	"net/http"
	httpconst "nutri-plans-api/constants/http"
	msgconst "nutri-plans-api/constants/message"
	"nutri-plans-api/usecases"
	httputil "nutri-plans-api/utils/http"
	tokenutil "nutri-plans-api/utils/token"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type mealController struct {
	mealUsecase usecases.MealUsecase

	tokenUtil tokenutil.TokenUtil
}

func NewMealController(
	mealUsecase usecases.MealUsecase,
	tokenUtil tokenutil.TokenUtil,
) *mealController {
	return &mealController{
		mealUsecase: mealUsecase,
		tokenUtil:   tokenUtil,
	}
}

func (m *mealController) GetTodayMeal(c echo.Context) error {
	claims := m.tokenUtil.GetClaims(c)

	res, err := m.mealUsecase.GetTodayMeal(c, claims.UID)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgMealNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgGetTodayMealFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgGetTodayMealFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(c, http.StatusOK, msgconst.MsgGetTodayMealSuccess, res)
}
