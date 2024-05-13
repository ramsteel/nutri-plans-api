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

func TestNewDietaryPreferenceTypeUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewDietaryPreferenceTypeUsecase(
			mockrepo.NewMockDietaryPreferenceTypeRepository(t),
		),
	)
}

func TestGetDietaryPreferenceTypes(t *testing.T) {
	example := &[]entities.DietaryPreferenceType{
		{
			ID:          1,
			Name:        "vegan",
			Description: "some description.",
			CreatedAt:   time.UnixMilli(1714757476909),
			UpdatedAt:   time.UnixMilli(1714757476909),
			DeletedAt:   gorm.DeletedAt{},
		},
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/dietary-preference-types", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()
	mockDietaryPreferenceTypeRepo := new(mockrepo.MockDietaryPreferenceTypeRepository)
	mockDietaryPreferenceTypeRepo.On("GetDietaryPreferenceTypes", ctx).Return(example, nil)
	countryUsecase := usecases.NewDietaryPreferenceTypeUsecase(mockDietaryPreferenceTypeRepo)
	dietaryPreferenceTypes, err := countryUsecase.GetDietaryPreferenceTypes(c)
	assert.NoError(t, err)
	assert.Equal(t, example, dietaryPreferenceTypes)
}

func TestUpdateDietaryPreferenceType(t *testing.T) {
	id := uint(1)

	r := &dto.DietaryPreferenceTypeRequest{
		Name:        "vegan",
		Description: "some description.",
	}

	d := &entities.DietaryPreferenceType{
		ID:          id,
		Name:        "vegan",
		Description: "some description.",
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/dietary-preference-types/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockDietaryPreferenceTypeRepo := new(mockrepo.MockDietaryPreferenceTypeRepository)
	mockDietaryPreferenceTypeRepo.On("UpdateDietaryPreferenceType", ctx, d).Return(nil)
	dietaryPrefUsecase := usecases.NewDietaryPreferenceTypeUsecase(mockDietaryPreferenceTypeRepo)

	err := dietaryPrefUsecase.UpdateDietaryPreferenceType(c, r, id)

	assert.NoError(t, err)
}

func TestCreateDietaryPreferenceType(t *testing.T) {
	r := &dto.DietaryPreferenceTypeRequest{
		Name:        "vegan",
		Description: "some description.",
	}

	d := &entities.DietaryPreferenceType{
		Name:        "vegan",
		Description: "some description.",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/dietary-preference-types", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockDietaryPreferenceTypeRepo := new(mockrepo.MockDietaryPreferenceTypeRepository)
	mockDietaryPreferenceTypeRepo.On("CreateDietaryPreferenceType", ctx, d).Return(nil)

	dietaryPrefUsecase := usecases.NewDietaryPreferenceTypeUsecase(mockDietaryPreferenceTypeRepo)
	err := dietaryPrefUsecase.CreateDietaryPreferenceType(c, r)
	assert.NoError(t, err)
}

func TestDeleteDietaryPreferenceType(t *testing.T) {
	id := uint(1)
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/dietary-preference-types/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()
	mockDietaryPreferenceTypeRepo := new(mockrepo.MockDietaryPreferenceTypeRepository)
	mockDietaryPreferenceTypeRepo.On("DeleteDietaryPreferenceType", ctx, id).Return(nil)
	dietaryPrefUsecase := usecases.NewDietaryPreferenceTypeUsecase(mockDietaryPreferenceTypeRepo)
	err := dietaryPrefUsecase.DeleteDietaryPreferenceType(c, id)
	assert.NoError(t, err)
}
