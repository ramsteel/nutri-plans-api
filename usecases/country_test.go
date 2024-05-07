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

func TestNewCountryUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewCountryUsecase(
			mockrepo.NewMockCountryRepository(t),
		),
	)
}

func TestGetCountries(t *testing.T) {
	example := &[]entities.Country{
		{
			ID:        1,
			Name:      "indonesia",
			CreatedAt: time.UnixMilli(1714757476909),
			UpdatedAt: time.UnixMilli(1714757476909),
			DeletedAt: gorm.DeletedAt{},
		},
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/countries", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()
	mockCountryRepo := new(mockrepo.MockCountryRepository)
	mockCountryRepo.On("GetAllCountries", ctx).Return(example, nil)
	countryUsecase := usecases.NewCountryUsecase(mockCountryRepo)
	countriess, err := countryUsecase.GetCountries(c)
	assert.NoError(t, err)
	assert.Equal(t, example, countriess)
}
