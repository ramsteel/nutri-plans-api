package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type DietaryPreferenceTypeRepository interface {
	GetDietaryPreferenceTypes(ctx context.Context) (*[]entities.DietaryPreferenceType, error)
}

type dietaryPreferenceTypeRepository struct {
	db *gorm.DB
}

func NewDietaryPreferenceTypeRepository(db *gorm.DB) *dietaryPreferenceTypeRepository {
	return &dietaryPreferenceTypeRepository{
		db: db,
	}
}

func (f *dietaryPreferenceTypeRepository) GetDietaryPreferenceTypes(ctx context.Context) (
	*[]entities.DietaryPreferenceType, error,
) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	dietaryPreferenceTypes := new([]entities.DietaryPreferenceType)
	if err := f.db.Find(&dietaryPreferenceTypes).Error; err != nil {
		return nil, err
	}

	return dietaryPreferenceTypes, nil
}
