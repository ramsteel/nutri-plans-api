package controllers

import (
	"context"
	"errors"
	"net/http"
	msgconst "nutri-plans-api/constants/message"
	statusconst "nutri-plans-api/constants/status"
	"nutri-plans-api/dto"
	"nutri-plans-api/usecases"
	errutil "nutri-plans-api/utils/error"
	valutil "nutri-plans-api/utils/validation"

	httpconst "nutri-plans-api/constants/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userController struct {
	userUsecase usecases.UserUsecase
	validator   *valutil.Validator
}

func NewUserController(userUsecase usecases.UserUsecase, v *valutil.Validator) *userController {
	return &userController{
		userUsecase: userUsecase,
		validator:   v,
	}
}

func (u *userController) Register(c echo.Context) error {
	req := new(dto.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, &dto.BaseResponse{
			Status:  statusconst.StatusFailed,
			Message: msgconst.MsgMismatchedDataType,
		})
	}

	if err := u.validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, &dto.BaseResponse{
			Status:  statusconst.StatusFailed,
			Message: msgconst.MsgInvalidRequestData,
		})
	}

	if err := u.userUsecase.Register(c, req); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			return c.JSON(httpconst.StatusClientCancelledRequest, &dto.BaseResponse{
				Status:  statusconst.StatusFailed,
				Message: msgconst.MsgUserCreationFailed,
			})
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.JSON(http.StatusNotFound, &dto.BaseResponse{
				Status:  statusconst.StatusFailed,
				Message: msgconst.MsgCountryNotFound,
			})
		case errors.Is(err, errutil.ErrFailedHashingPassword):
			return c.JSON(http.StatusInternalServerError, &dto.BaseResponse{
				Status:  statusconst.StatusFailed,
				Message: msgconst.MsgFailedHashingPassword,
			})
		case errors.Is(err, gorm.ErrDuplicatedKey):
			return c.JSON(http.StatusConflict, &dto.BaseResponse{
				Status:  statusconst.StatusFailed,
				Message: msgconst.MsgUsernameExist,
			})
		default:
			return c.JSON(http.StatusInternalServerError, &dto.BaseResponse{
				Status:  statusconst.StatusFailed,
				Message: msgconst.MsgUserCreationFailed,
			})
		}
	}

	return c.JSON(http.StatusOK, &dto.BaseResponse{
		Status:  statusconst.StatusSuccess,
		Message: msgconst.SuccessUserCreated,
	})
}
