package usecases

import (
	"context"
	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"

	"github.com/labstack/echo/v4"
)

type FoodTypeUsecase interface {
	GetFoodTypes(c echo.Context) (*[]entities.FoodType, error)
	CreateFoodType(c echo.Context, r *dto.FoodTypeRequest) error
	UpdateFoodType(c echo.Context, r *dto.FoodTypeRequest, id uint) error
	DeleteFoodType(c echo.Context, id uint) error
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

func (f *foodTypeUsecase) CreateFoodType(c echo.Context, r *dto.FoodTypeRequest) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	foodType := &entities.FoodType{Name: r.Name}

	return f.foodTypeRepo.CreateFoodType(ctx, foodType)
}

func (f *foodTypeUsecase) UpdateFoodType(c echo.Context, r *dto.FoodTypeRequest, id uint) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	foodType := &entities.FoodType{
		ID:   id,
		Name: r.Name,
	}

	return f.foodTypeRepo.UpdateFoodType(ctx, foodType)
}

func (f *foodTypeUsecase) DeleteFoodType(c echo.Context, id uint) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return f.foodTypeRepo.DeleteFoodType(ctx, id)
}
