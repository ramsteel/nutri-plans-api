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
	valutil "nutri-plans-api/utils/validation"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type adminController struct {
	adminUsecase usecases.AdminUsecase

	tokenUtil tokenutil.TokenUtil
	v         *valutil.Validator
}

func NewAdminController(
	adminUsecase usecases.AdminUsecase,
	tokenUtil tokenutil.TokenUtil,
	v *valutil.Validator,
) *adminController {
	return &adminController{
		adminUsecase: adminUsecase,
		tokenUtil:    tokenUtil,
		v:            v,
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
