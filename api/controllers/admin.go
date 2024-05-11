package controllers

import (
	"context"
	"errors"
	"net/http"
	httpconst "nutri-plans-api/constants/http"
	msgconst "nutri-plans-api/constants/message"
	"nutri-plans-api/dto"
	"nutri-plans-api/usecases"
	"strconv"

	httputil "nutri-plans-api/utils/http"
	tokenutil "nutri-plans-api/utils/token"
	valutil "nutri-plans-api/utils/validation"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type adminController struct {
	adminUsecase    usecases.AdminUsecase
	foodTypeUsecase usecases.FoodTypeUsecase

	tokenUtil tokenutil.TokenUtil
	v         *valutil.Validator
}

func NewAdminController(
	adminUsecase usecases.AdminUsecase,
	foodTypeUsecase usecases.FoodTypeUsecase,
	tokenUtil tokenutil.TokenUtil,
	v *valutil.Validator,
) *adminController {
	return &adminController{
		adminUsecase:    adminUsecase,
		foodTypeUsecase: foodTypeUsecase,
		tokenUtil:       tokenUtil,
		v:               v,
	}
}

func (a *adminController) GetAdminProfile(c echo.Context) error {
	claims := a.tokenUtil.GetClaims(c)

	res, err := a.adminUsecase.GetAdminProfile(c, claims.UID)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgAdminNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgGetAdminProfileFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgGetAdminProfileFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(c, http.StatusOK, msgconst.MsgGetAdminProfileSuccess, res)
}

func (a *adminController) CreateFoodType(c echo.Context) error {
	req := new(dto.FoodTypeRequest)
	if err := c.Bind(req); err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgMismatchedDataType,
		)
	}

	if err := a.v.Validate(req); err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	err := a.foodTypeUsecase.CreateFoodType(c, req)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgCreateFoodTypeFailed
		case errors.Is(err, gorm.ErrDuplicatedKey):
			code = http.StatusConflict
			msg = msgconst.MsgFoodTypeExist
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgCreateFoodTypeFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusCreated,
		msgconst.MsgCreateFoodTypeSuccess,
		nil,
	)
}

func (a *adminController) UpdateFoodType(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	req := new(dto.FoodTypeRequest)
	if err := c.Bind(req); err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgMismatchedDataType,
		)
	}

	if err := a.v.Validate(req); err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	err = a.foodTypeUsecase.UpdateFoodType(c, req, uint(intID))
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgFoodTypeNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgUpdateFoodTypeFailed
		case errors.Is(err, gorm.ErrDuplicatedKey):
			code = http.StatusConflict
			msg = msgconst.MsgFoodTypeExist
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgUpdateFoodTypeFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgUpdateFoodTypeSuccess,
		nil,
	)
}

func (a *adminController) DeleteFoodType(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	err = a.foodTypeUsecase.DeleteFoodType(c, uint(intID))
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgFoodTypeNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgDeleteFoodTypeFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgDeleteFoodTypeFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgDeleteFoodTypeSuccess,
		nil,
	)
}
