package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type MealTypeRepository interface {
	GetMealTypes(ctx context.Context) (*[]entities.MealType, error)
}

type mealTypeRepository struct {
	db *gorm.DB
}

func NewMealTypeRepository(db *gorm.DB) *mealTypeRepository {
	return &mealTypeRepository{
		db: db,
	}
}

func (f *mealTypeRepository) GetMealTypes(ctx context.Context) (*[]entities.MealType, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	mealTypes := new([]entities.MealType)
	if err := f.db.Find(&mealTypes).Error; err != nil {
		return nil, err
	}

	return mealTypes, nil
}
