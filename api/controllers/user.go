package controllers

import (
	"context"
	"errors"
	"net/http"
	msgconst "nutri-plans-api/constants/message"
	"nutri-plans-api/dto"
	"nutri-plans-api/usecases"
	errutil "nutri-plans-api/utils/error"
	httputil "nutri-plans-api/utils/http"
	tokenutil "nutri-plans-api/utils/token"
	valutil "nutri-plans-api/utils/validation"

	httpconst "nutri-plans-api/constants/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userController struct {
	userUsecase usecases.UserUsecase

	validator *valutil.Validator
	tokenUtil tokenutil.TokenUtil
}

func NewUserController(
	userUsecase usecases.UserUsecase,
	v *valutil.Validator,
	t tokenutil.TokenUtil,
) *userController {
	return &userController{
		userUsecase: userUsecase,
		validator:   v,
		tokenUtil:   t,
	}
}

func (u *userController) Register(c echo.Context) error {
	req := new(dto.RegisterRequest)
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

	if err := u.userUsecase.Register(c, req); err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgUserCreationFailed
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgCountryNotFound
		case errors.Is(err, errutil.ErrFailedHashingPassword):
			code = http.StatusInternalServerError
			msg = msgconst.MsgFailedHashingPassword
		case errors.Is(err, gorm.ErrDuplicatedKey):
			code = http.StatusConflict
			msg = msgconst.MsgUserExist
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgUserCreationFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(c, http.StatusCreated, msgconst.MsgUserCreated, nil)
}

func (u *userController) Login(c echo.Context) error {
	req := new(dto.LoginRequest)
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

	res, err := u.userUsecase.Login(c, req)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgLoginFailed
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgUnregisteredEmail
		case errors.Is(err, errutil.ErrPasswordMismatch):
			code = http.StatusUnauthorized
			msg = msgconst.MsgPasswordMismatch
		case errors.Is(err, errutil.ErrFailedGeneratingToken):
			code = http.StatusInternalServerError
			msg = msgconst.MsgFailedGeneratingToken
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgLoginFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(c, http.StatusOK, msgconst.MsgLoginSuccess, res)
}

func (u *userController) GetUserDetail(c echo.Context) error {
	claims := u.tokenUtil.GetClaims(c)

	user, err := u.userUsecase.GetUserByID(c, claims.UID)
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
			msg = msgconst.MsgGetUserFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgGetUserFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(c, http.StatusOK, msgconst.MsgRetrieveUserSuccess, user)
}

func (u *userController) UpdateUser(c echo.Context) error {
	claims := u.tokenUtil.GetClaims(c)

	req := new(dto.UpdateUserRequest)
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

	err := u.userUsecase.UpdateUser(c, claims.UID, req)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgUpdateUserFailed
		case errors.Is(err, gorm.ErrForeignKeyViolated):
			code = http.StatusNotFound
			msg = msgconst.MsgCountryNotFound
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgUpdateUserFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(c, http.StatusOK, msgconst.MsgUpdateUserSuccess, nil)
}
