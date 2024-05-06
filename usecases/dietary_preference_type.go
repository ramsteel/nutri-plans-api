package usecases

import (
	"context"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"

	"github.com/labstack/echo/v4"
)

type DietaryPreferenceTypeUsecase interface {
	GetDietaryPreferenceTypes(c echo.Context) (*[]entities.DietaryPreferenceType, error)
}

type dietaryPreferenceTypeUsecase struct {
	dietaryPreferenceTypeRepo repositories.DietaryPreferenceTypeRepository
}

func NewDietaryPreferenceTypeUsecase(
	dietaryPreferenceTypeRepo repositories.DietaryPreferenceTypeRepository,
) *dietaryPreferenceTypeUsecase {
	return &dietaryPreferenceTypeUsecase{
		dietaryPreferenceTypeRepo: dietaryPreferenceTypeRepo,
	}
}

func (f *dietaryPreferenceTypeUsecase) GetDietaryPreferenceTypes(c echo.Context) (
	*[]entities.DietaryPreferenceType, error,
) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return f.dietaryPreferenceTypeRepo.GetDietaryPreferenceTypes(ctx)
}
