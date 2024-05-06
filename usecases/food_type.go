package usecases

import (
	"context"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"

	"github.com/labstack/echo/v4"
)

type FoodTypeUsecase interface {
	GetFoodTypes(c echo.Context) (*[]entities.FoodType, error)
}

type foodTypeUsecase struct {
	foodTypeRepo repositories.FoodTypeRepository
}

func NewFoodTypeUsecase(foodTypeRepo repositories.FoodTypeRepository) *foodTypeUsecase {
	return &foodTypeUsecase{
		foodTypeRepo: foodTypeRepo,
	}
}

func (f *foodTypeUsecase) GetFoodTypes(c echo.Context) (*[]entities.FoodType, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return f.foodTypeRepo.GetFoodTypes(ctx)
}
