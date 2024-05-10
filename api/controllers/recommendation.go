package controllers

import (
	"context"
	"errors"
	"net/http"
	httpconst "nutri-plans-api/constants/http"
	msgconst "nutri-plans-api/constants/message"
	"nutri-plans-api/usecases"
	errutil "nutri-plans-api/utils/error"
	httputil "nutri-plans-api/utils/http"
	tokenutil "nutri-plans-api/utils/token"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type recommendationController struct {
	u usecases.RecommendationUsecase

	tokenUtil tokenutil.TokenUtil
}

func NewRecommendationController(
	u usecases.RecommendationUsecase,
	tokenUtil tokenutil.TokenUtil,
) *recommendationController {
	return &recommendationController{
		u:         u,
		tokenUtil: tokenUtil,
	}
}

func (r *recommendationController) GetRecommendation(c echo.Context) error {
	claims := r.tokenUtil.GetClaims(c)

	res, err := r.u.GetRecommendation(c, claims.UID)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgGetRecommendationFailed
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgUnregisteredUser
		case errors.Is(err, errutil.ErrExternalService):
			code = http.StatusBadGateway
			msg = msgconst.MsgExternalServiceError
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgGetRecommendationFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgGetRecommendationSuccess,
		res,
	)
}
