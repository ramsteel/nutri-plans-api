package usecases

import (
	"context"
	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserPreferenceUsecase interface {
	UpdateUserPreference(c echo.Context, id uuid.UUID, r *dto.UserPreferenceRequest) error
	GetUserPreference(c echo.Context, id uuid.UUID) (*entities.UserPreference, error)
}

type userPreferenceUsecase struct {
	userPreferenceRepo repositories.UserPreferenceRepository
	dietaryRestRepo    repositories.DietaryRestrictionRepository
}

func NewUserPreferenceUsecase(
	userPreferenceRepo repositories.UserPreferenceRepository,
	dietaryRestRepo repositories.DietaryRestrictionRepository,
) *userPreferenceUsecase {
	return &userPreferenceUsecase{
		userPreferenceRepo: userPreferenceRepo,
		dietaryRestRepo:    dietaryRestRepo,
	}
}

func (u *userPreferenceUsecase) UpdateUserPreference(
	c echo.Context,
	id uuid.UUID,
	r *dto.UserPreferenceRequest,
) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	if err := u.dietaryRestRepo.DeleteDietaryRestriction(ctx, id); err != nil {
		return err
	}

	dietaryRestrictions, err := u.dietaryRestRepo.GetDeletedDietaryRestrictions(ctx, id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if r.DietaryRestrictions != nil {
		filterDietaryRestrictions(dietaryRestrictions, r.DietaryRestrictions)
	}

	userPreference := &entities.UserPreference{
		UserID:                  id,
		FoodTypeID:              r.FoodTypeID,
		DrinkTypeID:             r.DrinkTypeID,
		DietaryPreferenceTypeID: r.DietaryPreferenceTypeID,
		DietaryRestrictions:     dietaryRestrictions,
	}

	return u.userPreferenceRepo.UpdateUserPreference(ctx, userPreference)
}

func (u *userPreferenceUsecase) GetUserPreference(
	c echo.Context,
	id uuid.UUID,
) (*entities.UserPreference, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return u.userPreferenceRepo.GetUserPreference(ctx, id)
}

func filterDietaryRestrictions(dietaryRestrictions *[]entities.DietaryRestriction, data *[]string) {
	nameFlagged := make(map[string]bool)
	lenDietaryRestrictions := len(*dietaryRestrictions)
	idx := 0
	for _, name := range *data {
		name = strings.ToLower(name)
		if _, ok := nameFlagged[name]; ok {
			continue
		}

		if idx <= lenDietaryRestrictions-1 && lenDietaryRestrictions > 0 {
			(*dietaryRestrictions)[idx].Name = name
			(*dietaryRestrictions)[idx].DeletedAt = gorm.DeletedAt{}
		} else {
			*dietaryRestrictions = append(*dietaryRestrictions, entities.DietaryRestriction{
				Name: name,
			})
		}

		nameFlagged[name] = true
		idx++
	}
}
