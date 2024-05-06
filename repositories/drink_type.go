package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type DrinkTypeRepository interface {
	GetDrinkTypes(ctx context.Context) (*[]entities.DrinkType, error)
}

type drinkTypeRepository struct {
	db *gorm.DB
}

func NewDrinkTypeRepository(db *gorm.DB) *drinkTypeRepository {
	return &drinkTypeRepository{
		db: db,
	}
}

func (f *drinkTypeRepository) GetDrinkTypes(ctx context.Context) (*[]entities.DrinkType, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	drinkTypes := new([]entities.DrinkType)
	if err := f.db.Find(&drinkTypes).Error; err != nil {
		return nil, err
	}

	return drinkTypes, nil
}
