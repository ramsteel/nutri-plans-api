package repositories

import (
	"context"
	"nutri-plans-api/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MealItemRepository interface {
	GetCalculatedNutrients(
		ctx context.Context,
		id uuid.UUID,
		start, end time.Time,
	) (*entities.CalculatedNutrients, error)
	GetMealItemByID(ctx context.Context, id uint64) (*entities.MealItem, error)
}

type mealItemRepository struct {
	db *gorm.DB
}

func NewMealItemRepository(db *gorm.DB) *mealItemRepository {
	return &mealItemRepository{
		db: db,
	}
}

func (m *mealItemRepository) GetCalculatedNutrients(
	ctx context.Context,
	id uuid.UUID,
	start, end time.Time,
) (*entities.CalculatedNutrients, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	calculatedNutrients := new(entities.CalculatedNutrients)
	err := m.db.Model(&entities.MealItem{}).
		Select(
			"SUM(calories) as total_calories, SUM(carbohydrate) as total_carbohydrate, SUM(protein) as total_protein, SUM(fat) as total_fat, SUM(cholesterol) as total_cholesterol, SUM(sugars) as total_sugars",
		).
		Where("meal_id = ? AND created_at BETWEEN ? AND ?", id, start, end).
		Group("meal_id").
		Find(calculatedNutrients).Error
	if err != nil {
		return nil, err
	}

	return calculatedNutrients, nil
}

func (m *mealItemRepository) GetMealItemByID(
	ctx context.Context,
	id uint64,
) (*entities.MealItem, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	mealItem := new(entities.MealItem)
	if err := m.db.First(mealItem, id).Error; err != nil {
		return nil, err
	}

	return mealItem, nil
}
