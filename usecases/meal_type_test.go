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

func TestNewMealTypeUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewMealTypeUsecase(
			mockrepo.NewMockMealTypeRepository(t),
		),
	)
}

func TestGetMealTypes(t *testing.T) {
	example := &[]entities.MealType{
		{
			ID:        uint(1),
			Name:      "breakfast",
			CreatedAt: time.UnixMilli(1714757476909),
			UpdatedAt: time.UnixMilli(1714757476909),
			DeletedAt: gorm.DeletedAt{},
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/meal-types", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockMealTypeRepo := new(mockrepo.MockMealTypeRepository)
	mockMealTypeRepo.On("GetMealTypes", ctx).Return(example, nil)

	countryUsecase := usecases.NewMealTypeUsecase(mockMealTypeRepo)
	mealTypes, err := countryUsecase.GetMealTypes(c)

	assert.NoError(t, err)
	assert.Equal(t, example, mealTypes)
}
