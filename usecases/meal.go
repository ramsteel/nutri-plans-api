package usecases

import (
	"context"
	"errors"
	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"
	dateutil "nutri-plans-api/utils/date"
	errutil "nutri-plans-api/utils/error"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MealUsecase interface {
	GetTodayMeal(c echo.Context, uid uuid.UUID) (*entities.Meal, error)
	AddMeal(c echo.Context, r *dto.MealItemRequest, uid uuid.UUID) error
	UpdateMeal(c echo.Context, r *dto.MealItemRequest, uid uuid.UUID, id uint64) error
	GetMealItemByID(c echo.Context, uid uuid.UUID, id uint64) (*entities.MealItem, error)
}

type mealUsecase struct {
	mealRepo     repositories.MealRepository
	mealItemRepo repositories.MealItemRepository
}

func NewMealUsecase(
	mealRepo repositories.MealRepository,
	mealItemRepo repositories.MealItemRepository,
) *mealUsecase {
	return &mealUsecase{
		mealRepo:     mealRepo,
		mealItemRepo: mealItemRepo,
	}
}

func (m *mealUsecase) GetTodayMeal(c echo.Context, uid uuid.UUID) (*entities.Meal, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	start, end := dateutil.GetTodayRange()

	return m.mealRepo.GetTodayMeal(ctx, uid, start, end)
}

func (m *mealUsecase) AddMeal(c echo.Context, r *dto.MealItemRequest, uid uuid.UUID) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	var meal *entities.Meal
	todayMeal, err := m.GetTodayMeal(c, uid)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) || len(todayMeal.MealItems) == 0 {
		meal = &entities.Meal{
			UserID: uid,
			MealItems: []entities.MealItem{
				{
					MealTypeID:   r.MealTypeID,
					ItemName:     r.ItemName,
					Qty:          r.Qty,
					Unit:         r.Unit,
					Weight:       r.Weight,
					Calories:     *r.Calories,
					Carbohydrate: *r.Carbohydrate,
					Protein:      *r.Protein,
					Fat:          *r.Fat,
					Cholesterol:  *r.Cholesterol,
					Sugars:       *r.Sugars,
				},
			},
			CalculatedNutrients: entities.CalculatedNutrients{
				TotalCalories:     *r.Calories,
				TotalCarbohydrate: *r.Carbohydrate,
				TotalProtein:      *r.Protein,
				TotalFat:          *r.Fat,
				TotalCholesterol:  *r.Cholesterol,
				TotalSugars:       *r.Sugars,
			},
		}
	} else {
		start, end := dateutil.GetTodayRange()
		calcNutrition, err := m.mealItemRepo.GetCalculatedNutrients(ctx, todayMeal.ID, start, end)
		if err != nil {
			return err
		}
		meal = &entities.Meal{
			ID:     todayMeal.ID,
			UserID: uid,
			MealItems: []entities.MealItem{
				{
					MealTypeID:   r.MealTypeID,
					ItemName:     r.ItemName,
					Qty:          r.Qty,
					Unit:         r.Unit,
					Weight:       r.Weight,
					Calories:     *r.Calories,
					Carbohydrate: *r.Carbohydrate,
					Protein:      *r.Protein,
					Fat:          *r.Fat,
					Cholesterol:  *r.Cholesterol,
					Sugars:       *r.Sugars,
				},
			},
			CalculatedNutrients: entities.CalculatedNutrients{
				TotalCalories:     *r.Calories + calcNutrition.TotalCalories,
				TotalCarbohydrate: *r.Carbohydrate + calcNutrition.TotalCarbohydrate,
				TotalProtein:      *r.Protein + calcNutrition.TotalProtein,
				TotalFat:          *r.Fat + calcNutrition.TotalFat,
				TotalCholesterol:  *r.Cholesterol + calcNutrition.TotalCholesterol,
				TotalSugars:       *r.Sugars + calcNutrition.TotalSugars,
			},
			CreatedAt: todayMeal.CreatedAt,
		}
	}

	return m.mealRepo.AddMeal(ctx, meal)
}

func (m *mealUsecase) UpdateMeal(
	c echo.Context,
	r *dto.MealItemRequest,
	uid uuid.UUID,
	id uint64,
) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	todayMeal, err := m.GetTodayMeal(c, uid)
	if err != nil {
		return err
	}

	if todayMeal.UserID != uid {
		return errutil.ErrForbiddenResource
	}

	mealItem, err := m.mealItemRepo.GetMealItemByID(ctx, id)
	if err != nil {
		return err
	}

	var (
		totalCalories     = todayMeal.TotalCalories + (*r.Calories - mealItem.Calories)
		totalCarbohydrate = todayMeal.TotalCarbohydrate + (*r.Carbohydrate - mealItem.Carbohydrate)
		totalProtein      = todayMeal.TotalProtein + (*r.Protein - mealItem.Protein)
		totalFat          = todayMeal.TotalFat + (*r.Fat - mealItem.Fat)
		totalCholesterol  = todayMeal.TotalCholesterol + (*r.Cholesterol - mealItem.Cholesterol)
		totalSugars       = todayMeal.TotalSugars + (*r.Sugars - mealItem.Sugars)
	)
	meal := &entities.Meal{
		ID:     todayMeal.ID,
		UserID: uid,
		MealItems: []entities.MealItem{
			{
				ID:           mealItem.ID,
				MealTypeID:   r.MealTypeID,
				ItemName:     r.ItemName,
				Qty:          r.Qty,
				Unit:         r.Unit,
				Weight:       r.Weight,
				Calories:     *r.Calories,
				Carbohydrate: *r.Carbohydrate,
				Protein:      *r.Protein,
				Fat:          *r.Fat,
				Cholesterol:  *r.Cholesterol,
				Sugars:       *r.Sugars,
			},
		},
		CalculatedNutrients: entities.CalculatedNutrients{
			TotalCalories:     totalCalories,
			TotalCarbohydrate: totalCarbohydrate,
			TotalProtein:      totalProtein,
			TotalFat:          totalFat,
			TotalCholesterol:  totalCholesterol,
			TotalSugars:       totalSugars,
		},
	}

	return m.mealRepo.UpdateMeal(ctx, meal)
}

func (m *mealUsecase) GetMealItemByID(
	c echo.Context,
	uid uuid.UUID,
	id uint64,
) (*entities.MealItem, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	todayMeal, err := m.GetTodayMeal(c, uid)
	if err != nil {
		return nil, err
	}

	if todayMeal.UserID != uid {
		return nil, errutil.ErrForbiddenResource
	}

	return m.mealItemRepo.GetMealItemByID(ctx, id)
}
