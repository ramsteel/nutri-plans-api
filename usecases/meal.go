package usecases

import (
	"context"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"
	dateutil "nutri-plans-api/utils/date"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type MealUsecase interface {
	GetTodayMeal(c echo.Context, uid uuid.UUID) (*entities.Meal, error)
}

type mealUsecase struct {
	mealRepo repositories.MealRepository
}

func NewMealUsecase(mealRepo repositories.MealRepository) *mealUsecase {
	return &mealUsecase{
		mealRepo: mealRepo,
	}
}

func (m *mealUsecase) GetTodayMeal(c echo.Context, uid uuid.UUID) (*entities.Meal, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	start, end := dateutil.GetTodayRange()

	return m.mealRepo.GetTodayMeal(ctx, uid, start, end)
}
