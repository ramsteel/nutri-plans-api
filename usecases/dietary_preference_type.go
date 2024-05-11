package usecases

import (
	"context"
	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"

	"github.com/labstack/echo/v4"
)

type DietaryPreferenceTypeUsecase interface {
	GetDietaryPreferenceTypes(c echo.Context) (*[]entities.DietaryPreferenceType, error)
	CreateDietaryPreferenceType(c echo.Context, r *dto.DietaryPreferenceTypeRequest) error
	UpdateDietaryPreferenceType(c echo.Context, r *dto.DietaryPreferenceTypeRequest, id uint) error
	DeleteDietaryPreferenceType(c echo.Context, id uint) error
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

func (f *dietaryPreferenceTypeUsecase) CreateDietaryPreferenceType(
	c echo.Context,
	r *dto.DietaryPreferenceTypeRequest,
) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	dietaryPreferenceType := &entities.DietaryPreferenceType{
		Name:        r.Name,
		Description: r.Description,
	}

	return f.dietaryPreferenceTypeRepo.CreateDietaryPreferenceType(ctx, dietaryPreferenceType)
}

func (f *dietaryPreferenceTypeUsecase) UpdateDietaryPreferenceType(
	c echo.Context,
	r *dto.DietaryPreferenceTypeRequest,
	id uint,
) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	dietaryPreferenceType := &entities.DietaryPreferenceType{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
	}

	return f.dietaryPreferenceTypeRepo.UpdateDietaryPreferenceType(ctx, dietaryPreferenceType)
}

func (f *dietaryPreferenceTypeUsecase) DeleteDietaryPreferenceType(c echo.Context, id uint) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return f.dietaryPreferenceTypeRepo.DeleteDietaryPreferenceType(ctx, id)
}
