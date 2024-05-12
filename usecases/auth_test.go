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

func TestNewAuthUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewAuthUsecase(
			mockrepo.NewMockAuthRepository(t),
		),
	)
}

func TestGetAllUsersAuth(t *testing.T) {
	id := uuid.New()
	example := &[]entities.Auth{
		{
			ID:       id,
			Username: "user",
			Email:    "user@example.com",
			RoleType: entities.RoleType{ID: 1, Name: "user"},
		},
	}

	mockAuthRepository := mockrepo.NewMockAuthRepository(t)
	usecase := usecases.NewAuthUsecase(mockAuthRepository)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/admin/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockAuthRepository.On("GetAllUsersAuths", ctx).Return(example, nil)

	res, err := usecase.GetAllUsersAuth(c)

	assert.NoError(t, err)
	assert.Equal(t, example, res)
	assert.Len(t, *res, 1)
}
