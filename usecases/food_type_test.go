package usecases_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	mockrepo "nutri-plans-api/mocks/repositories"
	"nutri-plans-api/usecases"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewFoodTypeUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewFoodTypeUsecase(
			mockrepo.NewMockFoodTypeRepository(t),
		),
	)
}

func TestGetFoodTypes(t *testing.T) {
	example := &[]entities.FoodType{
		{
			ID:        1,
			Name:      "dry",
			CreatedAt: time.UnixMilli(1714757476909),
			UpdatedAt: time.UnixMilli(1714757476909),
			DeletedAt: gorm.DeletedAt{},
		},
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/food-types", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()
	mockFoodTypeRepo := new(mockrepo.MockFoodTypeRepository)
	mockFoodTypeRepo.On("GetFoodTypes", ctx).Return(example, nil)
	countryUsecase := usecases.NewFoodTypeUsecase(mockFoodTypeRepo)
	foodTypes, err := countryUsecase.GetFoodTypes(c)
	assert.NoError(t, err)
	assert.Equal(t, example, foodTypes)
}

func TestCreateFoodType(t *testing.T) {
	r := &dto.FoodTypeRequest{
		Name: "dry",
	}

	d := &entities.FoodType{
		Name: "dry",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/food-types", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockFoodTypeRepo := new(mockrepo.MockFoodTypeRepository)
	mockFoodTypeRepo.On("CreateFoodType", ctx, d).Return(nil)

	countryUsecase := usecases.NewFoodTypeUsecase(mockFoodTypeRepo)
	err := countryUsecase.CreateFoodType(c, r)
	assert.NoError(t, err)
}

func TestUpdateFoodType(t *testing.T) {
	id := uint(1)

	r := &dto.FoodTypeRequest{
		Name: "dry",
	}

	d := &entities.FoodType{
		ID:   id,
		Name: "dry",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/food-types/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockFoodTypeRepo := new(mockrepo.MockFoodTypeRepository)
	mockFoodTypeRepo.On("UpdateFoodType", ctx, d).Return(nil)

	countryUsecase := usecases.NewFoodTypeUsecase(mockFoodTypeRepo)
	err := countryUsecase.UpdateFoodType(c, r, id)
	assert.NoError(t, err)
}

func TestDeleteFoodType(t *testing.T) {
	id := uint(1)
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/food-types/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()
	mockFoodTypeRepo := new(mockrepo.MockFoodTypeRepository)
	mockFoodTypeRepo.On("DeleteFoodType", ctx, id).Return(nil)
	drinkTypeUsecase := usecases.NewFoodTypeUsecase(mockFoodTypeRepo)
	err := drinkTypeUsecase.DeleteFoodType(c, id)
	assert.NoError(t, err)
}
