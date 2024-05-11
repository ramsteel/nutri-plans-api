package usecases

import (
	"context"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"

	"github.com/labstack/echo/v4"
)

type AuthUsecase interface {
	GetAllUsersAuth(c echo.Context) (*[]entities.Auth, error)
}

type authUsecase struct {
	authRepo repositories.AuthRepository
}

func NewAuthUsecase(authRepo repositories.AuthRepository) *authUsecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}

func (a *authUsecase) GetAllUsersAuth(c echo.Context) (*[]entities.Auth, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return a.authRepo.GetAllUsersAuths(ctx)
}
