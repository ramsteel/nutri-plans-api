package controllers

import (
	"context"
	"errors"
	"net/http"
	httpconst "nutri-plans-api/constants/http"
	msgconst "nutri-plans-api/constants/message"
	"nutri-plans-api/dto"
	"nutri-plans-api/usecases"
	httputil "nutri-plans-api/utils/http"
	tokenutil "nutri-plans-api/utils/token"
	valutil "nutri-plans-api/utils/validation"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userPreferenceController struct {
	userPreferenceUsecase usecases.UserPreferenceUsecase
	validator             *valutil.Validator
	tokenUtil             tokenutil.TokenUtil
}

func NewUserPreferenceController(
	userPreferenceUsecase usecases.UserPreferenceUsecase,
	v *valutil.Validator,
	t tokenutil.TokenUtil,
) *userPreferenceController {
	return &userPreferenceController{
		userPreferenceUsecase: userPreferenceUsecase,
		validator:             v,
		tokenUtil:             t,
	}
}

func (u *userPreferenceController) UpdateUserPreference(c echo.Context) error {
	claims := u.tokenUtil.GetClaims(c)

	req := new(dto.UserPreferenceRequest)
	if err := c.Bind(req); err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgMismatchedDataType,
		)
	}

	if err := u.validator.Validate(req); err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	err := u.userPreferenceUsecase.UpdateUserPreference(c, claims.UID, req)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, gorm.ErrForeignKeyViolated):
			code = http.StatusNotFound
			msg = msgconst.MsgPreferenceInputNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgUpdatePreferenceFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgUpdatePreferenceFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgPreferenceUpdated,
		nil,
	)
}

func (u *userPreferenceController) GetUserPreference(c echo.Context) error {
	claims := u.tokenUtil.GetClaims(c)

	res, err := u.userPreferenceUsecase.GetUserPreference(c, claims.UID)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgUnregisteredUser
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgGetPreferenceFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgGetPreferenceFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgRetrievePreferenceSuccess,
		res,
	)
}
