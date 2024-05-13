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

func TestNewDrinkTypeUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewDrinkTypeUsecase(
			mockrepo.NewMockDrinkTypeRepository(t),
		),
	)
}

func TestGetDrinkTypes(t *testing.T) {
	example := &[]entities.DrinkType{
		{
			ID:        1,
			Name:      "hot",
			CreatedAt: time.UnixMilli(1714757476909),
			UpdatedAt: time.UnixMilli(1714757476909),
			DeletedAt: gorm.DeletedAt{},
		},
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/drink-types", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()
	mockDrinkTypeRepo := new(mockrepo.MockDrinkTypeRepository)
	mockDrinkTypeRepo.On("GetDrinkTypes", ctx).Return(example, nil)
	countryUsecase := usecases.NewDrinkTypeUsecase(mockDrinkTypeRepo)
	drinkTypes, err := countryUsecase.GetDrinkTypes(c)
	assert.NoError(t, err)
	assert.Equal(t, example, drinkTypes)
}

func TestCreateDrinkType(t *testing.T) {

	r := &dto.DrinkTypeRequest{
		Name: "hot",
	}

	d := &entities.DrinkType{
		Name: "hot",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/drink-types", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()
	mockDrinkTypeRepo := new(mockrepo.MockDrinkTypeRepository)
	mockDrinkTypeRepo.On("CreateDrinkType", ctx, d).Return(nil)
	countryUsecase := usecases.NewDrinkTypeUsecase(mockDrinkTypeRepo)
	err := countryUsecase.CreateDrinkType(c, r)
	assert.NoError(t, err)
}

func TestUpdateDrinkType(t *testing.T) {
	id := uint(1)

	r := &dto.DrinkTypeRequest{
		Name: "hot",
	}

	d := &entities.DrinkType{
		ID:   id,
		Name: "hot",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/drink-types/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockDrinkTypeRepo := new(mockrepo.MockDrinkTypeRepository)
	mockDrinkTypeRepo.On("UpdateDrinkType", ctx, d).Return(nil)

	countryUsecase := usecases.NewDrinkTypeUsecase(mockDrinkTypeRepo)
	err := countryUsecase.UpdateDrinkType(c, r, id)
	assert.NoError(t, err)
}

func TestDeleteDrinkType(t *testing.T) {
	id := uint(1)
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/drink-types/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()
	mockDrinkTypeRepo := new(mockrepo.MockDrinkTypeRepository)
	mockDrinkTypeRepo.On("DeleteDrinkType", ctx, id).Return(nil)
	drinkTypeUsecase := usecases.NewDrinkTypeUsecase(mockDrinkTypeRepo)
	err := drinkTypeUsecase.DeleteDrinkType(c, id)
	assert.NoError(t, err)
}
