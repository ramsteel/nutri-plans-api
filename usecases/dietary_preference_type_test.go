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
