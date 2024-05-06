package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DietaryRestrictionRepository interface {
	DeleteDietaryRestriction(ctx context.Context, uid uuid.UUID) error
	GetDeletedDietaryRestrictions(
		ctx context.Context,
		userID uuid.UUID,
	) (*[]entities.DietaryRestriction, error)
}

type dietaryRestrictionRepository struct {
	db *gorm.DB
}

func NewDietaryRestrictionRepository(db *gorm.DB) *dietaryRestrictionRepository {
	return &dietaryRestrictionRepository{
		db: db,
	}
}

func (d *dietaryRestrictionRepository) DeleteDietaryRestriction(
	ctx context.Context,
	uid uuid.UUID,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	return d.db.Where("user_preference_id = ?", uid).Delete(&entities.DietaryRestriction{}).Error
}

func (d *dietaryRestrictionRepository) GetDeletedDietaryRestrictions(
	ctx context.Context,
	uid uuid.UUID,
) (*[]entities.DietaryRestriction, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	dietaryRestrictions := new([]entities.DietaryRestriction)
	err := d.db.Unscoped().Where("user_preference_id = ?", uid).Find(dietaryRestrictions).Error
	if err != nil {
		return nil, err
	}

	return dietaryRestrictions, nil
}
