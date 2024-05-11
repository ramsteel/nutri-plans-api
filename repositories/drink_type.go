package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type DrinkTypeRepository interface {
	GetDrinkTypes(ctx context.Context) (*[]entities.DrinkType, error)
	CreateDrinkType(ctx context.Context, drinkType *entities.DrinkType) error
	UpdateDrinkType(ctx context.Context, drinkType *entities.DrinkType) error
	DeleteDrinkType(ctx context.Context, id uint) error
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

func (f *drinkTypeRepository) CreateDrinkType(ctx context.Context, drinkType *entities.DrinkType) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	return f.db.Create(drinkType).Error
}

func (f *drinkTypeRepository) UpdateDrinkType(ctx context.Context, drinkType *entities.DrinkType) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	res := f.db.Updates(drinkType)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (f *drinkTypeRepository) DeleteDrinkType(ctx context.Context, id uint) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	res := f.db.Delete(&entities.DrinkType{}, id)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
