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
	adminUsecase                 usecases.AdminUsecase
	foodTypeUsecase              usecases.FoodTypeUsecase
	drinkTypeUsecase             usecases.DrinkTypeUsecase
	dietaryPreferenceTypeUsecase usecases.DietaryPreferenceTypeUsecase
	authUsecase                  usecases.AuthUsecase

	tokenUtil tokenutil.TokenUtil
	v         *valutil.Validator
}

func NewAdminController(
	adminUsecase usecases.AdminUsecase,
	foodTypeUsecase usecases.FoodTypeUsecase,
	drinkTypeUsecase usecases.DrinkTypeUsecase,
	dietaryPreferenceTypeUsecase usecases.DietaryPreferenceTypeUsecase,
	authUsecase usecases.AuthUsecase,
	tokenUtil tokenutil.TokenUtil,
	v *valutil.Validator,
) *adminController {
	return &adminController{
		adminUsecase:                 adminUsecase,
		foodTypeUsecase:              foodTypeUsecase,
		drinkTypeUsecase:             drinkTypeUsecase,
		dietaryPreferenceTypeUsecase: dietaryPreferenceTypeUsecase,
		authUsecase:                  authUsecase,
		tokenUtil:                    tokenUtil,
		v:                            v,
	}
}

// admin
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

// food types
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

// drink types
func (a *adminController) CreateDrinkType(c echo.Context) error {
	req := new(dto.DrinkTypeRequest)
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

	err := a.drinkTypeUsecase.CreateDrinkType(c, req)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgCreateDrinkTypeFailed
		case errors.Is(err, gorm.ErrDuplicatedKey):
			code = http.StatusConflict
			msg = msgconst.MsgDrinkTypeExist
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgCreateDrinkTypeFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusCreated,
		msgconst.MsgCreateDrinkTypeSuccess,
		nil,
	)
}

func (a *adminController) UpdateDrinkType(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	req := new(dto.DrinkTypeRequest)
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

	err = a.drinkTypeUsecase.UpdateDrinkType(c, req, uint(intID))
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgDrinkTypeNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgUpdateDrinkTypeFailed
		case errors.Is(err, gorm.ErrDuplicatedKey):
			code = http.StatusConflict
			msg = msgconst.MsgDrinkTypeExist
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgUpdateDrinkTypeFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgUpdateDrinkTypeSuccess,
		nil,
	)
}

func (a *adminController) DeleteDrinkType(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	err = a.drinkTypeUsecase.DeleteDrinkType(c, uint(intID))
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgDrinkTypeNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgDeleteDrinkTypeFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgDeleteDrinkTypeFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgDeleteDrinkTypeSuccess,
		nil,
	)
}

// dietary preference types
func (a *adminController) CreateDietaryPreferenceType(c echo.Context) error {
	req := new(dto.DietaryPreferenceTypeRequest)
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

	err := a.dietaryPreferenceTypeUsecase.CreateDietaryPreferenceType(c, req)
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgCreateDietaryPrefTypeFailed
		case errors.Is(err, gorm.ErrDuplicatedKey):
			code = http.StatusConflict
			msg = msgconst.MsgDietaryPrefTypeExist
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgCreateDietaryPrefTypeFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusCreated,
		msgconst.MsgCreateDietaryPrefTypeSuccess,
		nil,
	)
}

func (a *adminController) UpdateDietaryPreferenceType(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	req := new(dto.DietaryPreferenceTypeRequest)
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

	err = a.dietaryPreferenceTypeUsecase.UpdateDietaryPreferenceType(c, req, uint(intID))
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgDietaryPrefTypeNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgUpdateDietaryPrefTypeFailed
		case errors.Is(err, gorm.ErrDuplicatedKey):
			code = http.StatusConflict
			msg = msgconst.MsgDietaryPrefTypeExist
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgUpdateDietaryPrefTypeFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgUpdateDietaryPrefTypeSuccess,
		nil,
	)
}

func (a *adminController) DeleteDietaryPreferenceType(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return httputil.HandleErrorResponse(
			c,
			http.StatusBadRequest,
			msgconst.MsgInvalidRequestData,
		)
	}

	err = a.dietaryPreferenceTypeUsecase.DeleteDietaryPreferenceType(c, uint(intID))
	if err != nil {
		var (
			code int
			msg  string
		)

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			code = http.StatusNotFound
			msg = msgconst.MsgDietaryPrefTypeNotFound
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
			msg = msgconst.MsgDeleteDietaryPrefTypeFailed
		default:
			code = http.StatusInternalServerError
			msg = msgconst.MsgDeleteDietaryPrefTypeFailed
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(
		c,
		http.StatusOK,
		msgconst.MsgDeleteDietaryPrefTypeSuccess,
		nil,
	)
}

// users
func (a *adminController) GetAllUsersAuth(c echo.Context) error {
	res, err := a.authUsecase.GetAllUsersAuth(c)
	if err != nil {
		var (
			code int
			msg  string = msgconst.MsgGetAllUsersFailed
		)

		switch {
		case errors.Is(err, context.Canceled):
			code = httpconst.StatusClientCancelledRequest
		default:
			code = http.StatusInternalServerError
		}

		return httputil.HandleErrorResponse(c, code, msg)
	}

	return httputil.HandleSuccessResponse(c, http.StatusOK, msgconst.MsgGetAllUsersSuccess, res)
}
