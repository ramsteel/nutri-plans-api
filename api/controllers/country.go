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

type countryController struct {
	countryUsecase usecases.CountryUsecase
}

func NewCountryController(countryUsecase usecases.CountryUsecase) *countryController {
	return &countryController{
		countryUsecase: countryUsecase,
	}
}

func (cc *countryController) GetCountries(c echo.Context) error {
	res, err := cc.countryUsecase.GetAllCountry(c)
	if err != nil {
		var (
			code int    = http.StatusInternalServerError
			msg  string = msgconst.MsgGetCountriesFailed
		)

		if errors.Is(err, context.Canceled) {
			code = httpconst.StatusClientCancelledRequest
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(c, http.StatusOK, msgconst.MsgRetrieveCountrySuccess, res)
}
