package usecases

import (
	"context"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AdminUsecase interface {
	GetAdminProfile(c echo.Context, id uuid.UUID) (*entities.Admin, error)
}

type adminUsecase struct {
	adminRepo repositories.AdminRepository
}

func NewAdminUsecase(adminRepo repositories.AdminRepository) *adminUsecase {
	return &adminUsecase{
		adminRepo: adminRepo,
	}
}

func (a *adminUsecase) GetAdminProfile(c echo.Context, id uuid.UUID) (*entities.Admin, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return a.adminRepo.GetAdminProfile(ctx, id)
}
