package repositories

import (
	"context"
	"nutri-plans-api/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MealRepository interface {
	GetTodayMeal(ctx context.Context, uid uuid.UUID, start, end time.Time) (*entities.Meal, error)
	AddMeal(ctx context.Context, meal *entities.Meal) error
}

type mealRepository struct {
	db *gorm.DB
}

func NewMealRepository(db *gorm.DB) *mealRepository {
	return &mealRepository{
		db: db,
	}
}

func (m *mealRepository) GetTodayMeal(
	ctx context.Context,
	uid uuid.UUID,
	start, end time.Time,
) (*entities.Meal, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	meal := new(entities.Meal)
	err := m.db.Preload("MealItems.MealType").Preload(clause.Associations).Where(
		"user_id = ? AND created_at BETWEEN ? AND ?", uid, start, end,
	).First(meal).Error

	if err != nil {
		return nil, err
	}

	return meal, nil
}

func (m *mealRepository) AddMeal(ctx context.Context, meal *entities.Meal) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	return m.db.Save(meal).Error
}
