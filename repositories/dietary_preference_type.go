package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type DietaryPreferenceTypeRepository interface {
	GetDietaryPreferenceTypes(ctx context.Context) (*[]entities.DietaryPreferenceType, error)
	CreateDietaryPreferenceType(
		ctx context.Context,
		dietaryPreferenceType *entities.DietaryPreferenceType,
	) error
	UpdateDietaryPreferenceType(
		ctx context.Context,
		dietaryPreferenceType *entities.DietaryPreferenceType,
	) error
	DeleteDietaryPreferenceType(ctx context.Context, id uint) error
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

func (f *dietaryPreferenceTypeRepository) CreateDietaryPreferenceType(
	ctx context.Context,
	dietaryPreferenceType *entities.DietaryPreferenceType,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	return f.db.Create(dietaryPreferenceType).Error
}

func (f *dietaryPreferenceTypeRepository) UpdateDietaryPreferenceType(
	ctx context.Context,
	dietaryPreferenceType *entities.DietaryPreferenceType,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	res := f.db.Updates(dietaryPreferenceType)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (f *dietaryPreferenceTypeRepository) DeleteDietaryPreferenceType(
	ctx context.Context,
	id uint,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	res := f.db.Delete(&entities.DietaryPreferenceType{}, id)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
