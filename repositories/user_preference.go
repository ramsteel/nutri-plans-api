package repositories

import (
	"context"
	"nutri-plans-api/entities"
	"nutri-plans-api/utils/structconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserPreferenceRepository interface {
	CreateUserPreference(ctx context.Context, userPreference *entities.UserPreference) error
	UpdateUserPreference(ctx context.Context, userPreference *entities.UserPreference) error
	GetUserPreference(ctx context.Context, id uuid.UUID) (*entities.UserPreference, error)
	GetAllUserPreferences(ctx context.Context) (*[]entities.UserPreference, error)
}

type userPreferenceRepository struct {
	db *gorm.DB
}

func NewUserPreferenceRepository(db *gorm.DB) *userPreferenceRepository {
	return &userPreferenceRepository{
		db: db,
	}
}

func (u *userPreferenceRepository) CreateUserPreference(
	ctx context.Context,
	userPreference *entities.UserPreference,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	return u.db.Create(userPreference).Error
}

func (u *userPreferenceRepository) UpdateUserPreference(
	ctx context.Context,
	userPreference *entities.UserPreference,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	mainTableMap := structconv.ToMap(*userPreference)

	return u.db.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(userPreference).
		Updates(mainTableMap).Error
}

func (u *userPreferenceRepository) GetUserPreference(
	ctx context.Context,
	id uuid.UUID,
) (*entities.UserPreference, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	userPreference := new(entities.UserPreference)
	err := u.db.Preload(clause.Associations).First(userPreference, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return userPreference, nil
}

func (u *userPreferenceRepository) GetAllUserPreferences(
	ctx context.Context,
) (*[]entities.UserPreference, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	userPreferences := new([]entities.UserPreference)
	if err := u.db.Preload(clause.Associations).Find(userPreferences).Error; err != nil {
		return nil, err
	}

	return userPreferences, nil
}
