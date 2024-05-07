package usecases_test

import (
	"context"
	"net/http"
	"net/http/httptest"
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
