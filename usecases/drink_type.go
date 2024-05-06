package usecases

import (
	"context"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"

	"github.com/labstack/echo/v4"
)

type DrinkTypeUsecase interface {
	GetDrinkTypes(c echo.Context) (*[]entities.DrinkType, error)
}

type drinkTypeUsecase struct {
	drinkTypeRepo repositories.DrinkTypeRepository
}

func NewDrinkTypeUsecase(drinkTypeRepo repositories.DrinkTypeRepository) *drinkTypeUsecase {
	return &drinkTypeUsecase{
		drinkTypeRepo: drinkTypeRepo,
	}
}

func (f *drinkTypeUsecase) GetDrinkTypes(c echo.Context) (*[]entities.DrinkType, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return f.drinkTypeRepo.GetDrinkTypes(ctx)
}
