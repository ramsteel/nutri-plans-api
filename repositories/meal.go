package repositories

import (
	"context"
	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	structconvutil "nutri-plans-api/utils/structconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MealRepository interface {
	GetTodayMeal(ctx context.Context, uid uuid.UUID, start, end time.Time) (*entities.Meal, error)
	AddMeal(ctx context.Context, meal *entities.Meal) error
	UpdateMeal(ctx context.Context, meal *entities.Meal) error
	GetUserMeals(
		ctx context.Context,
		uid uuid.UUID,
		p *dto.PaginationRequest,
	) (*[]entities.Meal, int64, error)
	GetMealByID(ctx context.Context, uid uuid.UUID, id uuid.UUID) (*entities.Meal, error)
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

func (m *mealRepository) UpdateMeal(ctx context.Context, meal *entities.Meal) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	mapCalcNutrients := structconvutil.ToMap(meal.CalculatedNutrients)

	return m.db.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(meal).
		Updates(mapCalcNutrients).Error
}

func (m *mealRepository) GetUserMeals(
	ctx context.Context,
	uid uuid.UUID,
	p *dto.PaginationRequest,
) (*[]entities.Meal, int64, error) {
	if err := ctx.Err(); err != nil {
		return nil, 0, err
	}

	offset := (p.Page - 1) * p.Limit

	meals := new([]entities.Meal)
	tx := m.selectQuery(p.From, p.To)
	res := tx.Limit(p.Limit).
		Offset(offset).
		Find(meals, "user_id = ?", uid).
		Offset(-1).
		Find(&[]entities.Meal{})
	if res.Error != nil {
		return nil, 0, res.Error
	}

	return meals, res.RowsAffected, nil
}

func (m *mealRepository) GetMealByID(
	ctx context.Context,
	uid uuid.UUID,
	id uuid.UUID,
) (*entities.Meal, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	meal := new(entities.Meal)
	err := m.db.Preload("MealItems.MealType").
		Preload(clause.Associations).
		Where("user_id = ? AND id = ?", uid, id).
		First(meal).Error
	if err != nil {
		return nil, err
	}

	return meal, nil
}

func (m *mealRepository) selectQuery(from, to *time.Time) *gorm.DB {
	tx := m.db.Preload("MealItems.MealType").Preload(clause.Associations)
	if from == nil && to == nil {
		return tx
	} else if from == nil {
		return tx.Where("created_at <= ?", to)
	} else if to == nil {
		return tx.Where("created_at >= ?", from)
	} else {
		return tx.Where("created_at BETWEEN ? AND ?", from, to)
	}
}
