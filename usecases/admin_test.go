package usecases_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"nutri-plans-api/entities"
	mockrepo "nutri-plans-api/mocks/repositories"
	"nutri-plans-api/usecases"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewAdminUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewAdminUsecase(
			mockrepo.NewMockAdminRepository(t),
		),
	)
}

func TestGetAdminProfile(t *testing.T) {
	id := uuid.New()
	example := &entities.Admin{
		Auth: entities.Auth{
			ID:       id,
			Username: "admin",
			Email:    "admin@example.com",
			RoleType: entities.RoleType{ID: 2, Name: "admin"},
		},
	}
	mockAdminRepo := new(mockrepo.MockAdminRepository)
	u := usecases.NewAdminUsecase(mockAdminRepo)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockAdminRepo.On("GetAdminProfile", ctx, id).Return(example, nil)

	res, err := u.GetAdminProfile(c, id)

	assert.NoError(t, err)
	assert.Equal(t, example, res)
}
