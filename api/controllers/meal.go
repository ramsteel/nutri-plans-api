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
	tokenutil "nutri-plans-api/utils/token"
	valutil "nutri-plans-api/utils/validation"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type mealController struct {
	mealUsecase usecases.MealUsecase

	tokenUtil tokenutil.TokenUtil
	v         *valutil.Validator
}

func NewMealController(
	mealUsecase usecases.MealUsecase,
	tokenUtil tokenutil.TokenUtil,
	v *valutil.Validator,
) *mealController {
	return &mealController{
		mealUsecase: mealUsecase,
		tokenUtil:   tokenUtil,
		v:           v,
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

func (m *mealController) AddItemToMeal(c echo.Context) error {
	claims := m.tokenUtil.GetClaims(c)

	r := new(dto.MealItemRequest)
	if err := c.Bind(r); err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgMismatchedDataType,
		)
	}

	if err := m.v.Validate(r); err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	err := m.mealUsecase.AddMeal(c, r, claims.UID)
	if err != nil {

		if errors.Is(err, context.Canceled) {
			return httputil.HandleErrorResponse(
				c,
				httpconst.StatusClientCancelledRequest,
				msgconst.MsgAddItemToMealFailed,
			)
		}

		return httputil.HandleErrorResponse(
			c,
			http.StatusInternalServerError,
			msgconst.MsgAddItemToMealFailed,
		)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusCreated,
		msgconst.MsgAddItemToMealSuccess,
		nil,
	)
}

func (m *mealController) UpdateMealItem(c echo.Context) error {
	claims := m.tokenUtil.GetClaims(c)

	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	r := new(dto.MealItemRequest)
	if err := c.Bind(r); err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgMismatchedDataType,
		)
	}

	if err := m.v.Validate(r); err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	err = m.mealUsecase.UpdateMeal(c, r, claims.UID, uint64(intID))
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, errutil.ErrForbiddenResource):
			code = http.StatusForbidden
			msg = msgconst.MsgForbiddenResource
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgMealNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgUpdateMealItemFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgUpdateMealItemFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(c, http.StatusOK, msgconst.MsgUpdateMealItemSuccess, nil)
}

func (m *mealController) GetMealItemByID(c echo.Context) error {
	claims := m.tokenUtil.GetClaims(c)

	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	item, err := m.mealUsecase.GetMealItemByID(c, claims.UID, uint64(intID))
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, errutil.ErrForbiddenResource):
			code = http.StatusForbidden
			msg = msgconst.MsgForbiddenResource
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgMealItemNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgGetMealItemFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgGetMealItemFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(c, http.StatusOK, msgconst.MsgGetMealItemSuccess, item)
}

func (m *mealController) DeleteMealItem(c echo.Context) error {
	claims := m.tokenUtil.GetClaims(c)

	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	err = m.mealUsecase.DeleteMealItem(c, claims.UID, uint64(intID))
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, errutil.ErrForbiddenResource):
			code = http.StatusForbidden
			msg = msgconst.MsgForbiddenResource
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgMealItemNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgDeleteMealItemFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgDeleteMealItemFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(c, http.StatusOK, msgconst.MsgDeleteMealItemSuccess, nil)
}
