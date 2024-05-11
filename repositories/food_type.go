package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type FoodTypeRepository interface {
	GetFoodTypes(ctx context.Context) (*[]entities.FoodType, error)
	CreateFoodType(ctx context.Context, foodType *entities.FoodType) error
	UpdateFoodType(ctx context.Context, foodType *entities.FoodType) error
	DeleteFoodType(ctx context.Context, id uint) error
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

func (f *foodTypeRepository) CreateFoodType(ctx context.Context, foodType *entities.FoodType) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	return f.db.Create(foodType).Error
}

func (f *foodTypeRepository) UpdateFoodType(ctx context.Context, foodType *entities.FoodType) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	res := f.db.Updates(foodType)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (f *foodTypeRepository) DeleteFoodType(ctx context.Context, id uint) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	res := f.db.Delete(&entities.FoodType{}, id)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
