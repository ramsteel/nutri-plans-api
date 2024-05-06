package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type FoodTypeRepository interface {
	GetFoodTypes(ctx context.Context) (*[]entities.FoodType, error)
}

type foodTypeRepository struct {
	db *gorm.DB
}

func NewFoodTypeRepository(db *gorm.DB) *foodTypeRepository {
	return &foodTypeRepository{
		db: db,
	}
}

func (f *foodTypeRepository) GetFoodTypes(ctx context.Context) (*[]entities.FoodType, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	foodTypes := new([]entities.FoodType)
	if err := f.db.Find(&foodTypes).Error; err != nil {
		return nil, err
	}

	return foodTypes, nil
}
