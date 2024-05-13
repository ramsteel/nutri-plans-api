package usecases

import (
	"context"
	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"

	"github.com/labstack/echo/v4"
)

type DrinkTypeUsecase interface {
	GetDrinkTypes(c echo.Context) (*[]entities.DrinkType, error)
	CreateDrinkType(c echo.Context, r *dto.DrinkTypeRequest) error
	UpdateDrinkType(c echo.Context, r *dto.DrinkTypeRequest, id uint) error
	DeleteDrinkType(c echo.Context, id uint) error
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

func (f *drinkTypeUsecase) CreateDrinkType(c echo.Context, r *dto.DrinkTypeRequest) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	drinkType := &entities.DrinkType{Name: r.Name}

	return f.drinkTypeRepo.CreateDrinkType(ctx, drinkType)
}

func (f *drinkTypeUsecase) UpdateDrinkType(c echo.Context, r *dto.DrinkTypeRequest, id uint) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	drinkType := &entities.DrinkType{
		ID:   id,
		Name: r.Name,
	}

	return f.drinkTypeRepo.UpdateDrinkType(ctx, drinkType)
}

func (f *drinkTypeUsecase) DeleteDrinkType(c echo.Context, id uint) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return f.drinkTypeRepo.DeleteDrinkType(ctx, id)
}
