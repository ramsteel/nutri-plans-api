package usecases

import (
	"context"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"

	"github.com/labstack/echo/v4"
)

type MealTypeUsecase interface {
	GetMealTypes(c echo.Context) (*[]entities.MealType, error)
}

type mealTypeUsecase struct {
	mealTypeRepo repositories.MealTypeRepository
}

func NewMealTypeUsecase(mealTypeRepo repositories.MealTypeRepository) *mealTypeUsecase {
	return &mealTypeUsecase{
		mealTypeRepo: mealTypeRepo,
	}
}

func (f *mealTypeUsecase) GetMealTypes(c echo.Context) (*[]entities.MealType, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return f.mealTypeRepo.GetMealTypes(ctx)
}
