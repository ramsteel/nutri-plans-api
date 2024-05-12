package usecases_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	mockrepo "nutri-plans-api/mocks/repositories"
	"nutri-plans-api/usecases"
	dateutil "nutri-plans-api/utils/date"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type testCaseMeal struct {
	testCase
	todayMealID uuid.UUID
}

type testCasePaginationMeal struct {
	testCase
	totalData int
}

func TestNewMealUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewMealUsecase(
			mockrepo.NewMockMealRepository(t),
			mockrepo.NewMockMealItemRepository(t),
		),
	)
}

func TestGetTodayMeal(t *testing.T) {
	id := uuid.New()
	uid := uuid.New()
	meals := &entities.Meal{
		ID:     id,
		UserID: uid,
		MealItems: []entities.MealItem{{
			ID: uint64(1),
			MealType: entities.MealType{
				ID:        uint(1),
				Name:      "breakfast",
				CreatedAt: time.UnixMilli(1714757476909),
				UpdatedAt: time.UnixMilli(1714757476909),
				DeletedAt: gorm.DeletedAt{},
			},
			ItemName:     "egg",
			Qty:          1,
			Weight:       100,
			Unit:         "g",
			Calories:     70,
			Carbohydrate: 70,
			Protein:      70,
			Fat:          70,
			Cholesterol:  70,
			Sugars:       70,
			CreatedAt:    time.UnixMilli(1714757476909),
			UpdatedAt:    time.UnixMilli(1714757476909),
			DeletedAt:    gorm.DeletedAt{},
		}},
		CalculatedNutrients: entities.CalculatedNutrients{
			TotalCalories:     70,
			TotalCarbohydrate: 70,
			TotalProtein:      70,
			TotalFat:          70,
			TotalCholesterol:  70,
			TotalSugars:       70,
		},
		CreatedAt: time.UnixMilli(1714757476909),
		UpdatedAt: time.UnixMilli(1714757476909),
		DeletedAt: gorm.DeletedAt{},
	}

	start, end := dateutil.GetTodayRange()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/meals/today", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockMealRepo := new(mockrepo.MockMealRepository)
	mockMealRepo.On("GetTodayMeal", ctx, uid, start, end).Return(meals, nil)
	mockMealItemRepo := new(mockrepo.MockMealItemRepository)

	mealUsecase := usecases.NewMealUsecase(mockMealRepo, mockMealItemRepo)

	res, err := mealUsecase.GetTodayMeal(c, uid)
	assert.NoError(t, err)
	assert.Equal(t, meals, res)
}

func TestAddMeal(t *testing.T) {
	testCases := []testCase{
		{
			name: "success",
			errs: []error{nil, nil},
		},
		{
			name: "error today meal",
			errs: []error{errors.New("error today meal"), nil},
		},
		{
			name: "error add meal item",
			errs: []error{gorm.ErrRecordNotFound, nil},
		},
		{
			name: "error get calculated",
			errs: []error{nil, errors.New("error get calculated")},
		},
	}

	var constNutrient float32 = 70

	id := uuid.New()
	uid := uuid.New()

	r := &dto.MealItemRequest{
		MealTypeID:   uint(1),
		ItemName:     "egg",
		Qty:          1,
		Weight:       100,
		Unit:         "g",
		Calories:     &constNutrient,
		Carbohydrate: &constNutrient,
		Protein:      &constNutrient,
		Fat:          &constNutrient,
		Cholesterol:  &constNutrient,
		Sugars:       &constNutrient,
	}

	calcNutrition := &entities.CalculatedNutrients{
		TotalCalories:     0,
		TotalCarbohydrate: 0,
		TotalProtein:      0,
		TotalFat:          0,
		TotalCholesterol:  0,
		TotalSugars:       0,
	}

	todayMeal := &entities.Meal{
		ID:     id,
		UserID: uid,
		MealItems: []entities.MealItem{{
			ID: uint64(1),
			MealType: entities.MealType{
				ID:        uint(1),
				Name:      "breakfast",
				CreatedAt: time.UnixMilli(1714757476909),
				UpdatedAt: time.UnixMilli(1714757476909),
				DeletedAt: gorm.DeletedAt{},
			},
			ItemName:     "egg",
			Qty:          1,
			Weight:       100,
			Unit:         "g",
			Calories:     70,
			Carbohydrate: 70,
			Protein:      70,
			Fat:          70,
			Cholesterol:  70,
			Sugars:       70,
			CreatedAt:    time.UnixMilli(1714757476909),
			UpdatedAt:    time.UnixMilli(1714757476909),
			DeletedAt:    gorm.DeletedAt{},
		}},
		CalculatedNutrients: entities.CalculatedNutrients{
			TotalCalories:     70,
			TotalCarbohydrate: 70,
			TotalProtein:      70,
			TotalFat:          70,
			TotalCholesterol:  70,
			TotalSugars:       70,
		},
		CreatedAt: time.UnixMilli(1714757476909),
		UpdatedAt: time.UnixMilli(1714757476909),
		DeletedAt: gorm.DeletedAt{},
	}

	start, end := dateutil.GetTodayRange()

	for idx, tc := range testCases {
		meal := &entities.Meal{
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
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/meals/items", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ctx, cancel := context.WithCancel(c.Request().Context())
		defer cancel()

		t.Run(tc.name, func(t *testing.T) {
			mockMealRepo := new(mockrepo.MockMealRepository)
			mockMealItemRepo := new(mockrepo.MockMealItemRepository)
			mockMealRepo.On("GetTodayMeal", ctx, uid, start, end).Return(todayMeal, tc.errs[0])
			mockMealItemRepo.On(
				"GetCalculatedNutrients",
				ctx,
				id,
				start,
				end,
			).Return(calcNutrition, tc.errs[1])

			if idx != 2 {
				meal.ID = id
				meal.CreatedAt = todayMeal.CreatedAt
			}

			mockMealRepo.On("AddMeal", ctx, meal).Return(nil)

			mealUsecase := usecases.NewMealUsecase(mockMealRepo, mockMealItemRepo)

			err := mealUsecase.AddMeal(c, r, uid)
			if idx != 0 && idx != 2 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestUpdateMeal(t *testing.T) {
	mealID := uuid.New()
	id := uint64(1)
	uid := uuid.New()
	testCases := []testCaseMeal{
		{
			testCase: testCase{
				name: "success",
				errs: []error{nil, nil},
			},
			todayMealID: mealID,
		},
		{
			testCase: testCase{
				name: "fail get today meal",
				errs: []error{errors.New("fail get today meal"), nil},
			},
			todayMealID: mealID,
		},
		{
			testCase: testCase{
				name: "fail get meal item",
				errs: []error{nil, errors.New("fail get meal item")},
			},
			todayMealID: mealID,
		},
		{
			testCase: testCase{
				name: "fail forbidden",
				errs: []error{nil, nil},
			},
			todayMealID: uuid.New(),
		},
	}

	nutrition := float32(70)

	r := &dto.MealItemRequest{
		MealTypeID:   1,
		ItemName:     "egg",
		Qty:          1,
		Unit:         "g",
		Weight:       100,
		Calories:     &nutrition,
		Carbohydrate: &nutrition,
		Protein:      &nutrition,
		Fat:          &nutrition,
		Cholesterol:  &nutrition,
		Sugars:       &nutrition,
	}

	mealItem := &entities.MealItem{
		ID:           id,
		MealTypeID:   r.MealTypeID,
		MealID:       mealID,
		ItemName:     r.ItemName,
		Qty:          r.Qty,
		Unit:         r.Unit,
		Weight:       r.Weight,
		Calories:     nutrition,
		Carbohydrate: nutrition,
		Protein:      nutrition,
		Fat:          nutrition,
		Cholesterol:  nutrition,
		Sugars:       nutrition,
	}

	start, end := dateutil.GetTodayRange()

	for idx, tc := range testCases {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/meals/items/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ctx, cancel := context.WithCancel(c.Request().Context())
		defer cancel()

		t.Run(tc.name, func(t *testing.T) {
			todayMeal := &entities.Meal{
				ID:        tc.todayMealID,
				UserID:    uid,
				MealItems: []entities.MealItem{},
				CalculatedNutrients: entities.CalculatedNutrients{
					TotalCalories:     0,
					TotalCarbohydrate: 0,
					TotalProtein:      0,
					TotalFat:          0,
				},
			}
			mockMealRepo := new(mockrepo.MockMealRepository)
			mockMealItemRepo := new(mockrepo.MockMealItemRepository)
			mockMealRepo.On("GetTodayMeal", ctx, uid, start, end).Return(todayMeal, tc.errs[0])
			mockMealItemRepo.On(
				"GetMealItemByID",
				ctx,
				id,
			).Return(mealItem, tc.errs[1])

			meal := &entities.Meal{
				ID:     mealID,
				UserID: uid,
				MealItems: []entities.MealItem{
					{
						ID:           id,
						MealTypeID:   r.MealTypeID,
						ItemName:     r.ItemName,
						Qty:          r.Qty,
						Unit:         r.Unit,
						Weight:       r.Weight,
						Calories:     *r.Calories,
						Protein:      *r.Protein,
						Fat:          *r.Fat,
						Sugars:       *r.Sugars,
						Cholesterol:  *r.Cholesterol,
						Carbohydrate: *r.Carbohydrate,
					},
				},
				CalculatedNutrients: entities.CalculatedNutrients{
					TotalCalories:     0,
					TotalCarbohydrate: 0,
					TotalProtein:      0,
					TotalFat:          0,
				},
			}
			mockMealRepo.On("UpdateMeal", ctx, meal).Return(nil)

			mealUsecase := usecases.NewMealUsecase(mockMealRepo, mockMealItemRepo)

			err := mealUsecase.UpdateMeal(c, r, uid, id)

			if idx == 0 {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}

}

func TestGetMealItemByID(t *testing.T) {
	mealItemID := uint64(1)
	id := uuid.New()
	uid := uuid.New()

	testCases := []testCaseMeal{
		{
			testCase: testCase{
				name: "success",
				errs: []error{nil, nil},
			},
			todayMealID: id,
		},
		{
			testCase: testCase{
				name: "error today meal",
				errs: []error{errors.New("error today meal"), nil},
			},
			todayMealID: id,
		},
		{
			testCase: testCase{
				name: "error get meal item",
				errs: []error{nil, errors.New("error get meal item")},
			},
			todayMealID: id,
		},
		{
			testCase: testCase{
				name: "forbidden",
				errs: []error{nil, nil},
			},
			todayMealID: uuid.New(),
		},
	}

	mealItem := &entities.MealItem{
		ID:           mealItemID,
		MealID:       id,
		MealTypeID:   1,
		ItemName:     "test",
		Qty:          1,
		Unit:         "test",
		Weight:       1,
		Calories:     1,
		Protein:      1,
		Fat:          1,
		Sugars:       1,
		Cholesterol:  1,
		Carbohydrate: 1,
	}

	start, end := dateutil.GetTodayRange()

	for idx, tc := range testCases {
		todayMeal := &entities.Meal{
			ID:        tc.todayMealID,
			UserID:    uid,
			MealItems: []entities.MealItem{},
			CalculatedNutrients: entities.CalculatedNutrients{
				TotalCalories:     0,
				TotalCarbohydrate: 0,
				TotalProtein:      0,
				TotalFat:          0,
			},
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/meals/items/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ctx, cancel := context.WithCancel(c.Request().Context())
		defer cancel()

		mockMealRepo := new(mockrepo.MockMealRepository)
		mockMealItemRepo := new(mockrepo.MockMealItemRepository)
		mockMealRepo.On("GetTodayMeal", ctx, uid, start, end).Return(todayMeal, tc.errs[0])
		mockMealItemRepo.On(
			"GetMealItemByID",
			ctx,
			mealItemID,
		).Return(mealItem, tc.errs[1])

		mealUsecase := usecases.NewMealUsecase(mockMealRepo, mockMealItemRepo)

		_, err := mealUsecase.GetMealItemByID(c, uid, mealItemID)

		if idx == 0 {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}

}

func TestDeleteMealItem(t *testing.T) {
	uid := uuid.New()
	id := uint64(1)
	mealID := uuid.New()

	testCases := []testCaseMeal{
		{
			testCase: testCase{
				name: "success",
				errs: []error{nil, nil, nil},
			},
			todayMealID: mealID,
		},
		{
			testCase: testCase{
				name: "today meal error",
				errs: []error{errors.New("error today meal"), nil, nil},
			},
			todayMealID: mealID,
		},
		{
			testCase: testCase{
				name: "delete meal item error",
				errs: []error{nil, errors.New("delete meal item error"), nil},
			},
			todayMealID: mealID,
		},
		{
			testCase: testCase{
				name: "forbidden",
				errs: []error{nil, nil, nil},
			},
			todayMealID: uuid.New(),
		},
		{
			testCase: testCase{
				name: "error update meal",
				errs: []error{nil, nil, errors.New("error")},
			},
			todayMealID: mealID,
		},
	}

	start, end := dateutil.GetTodayRange()

	db, _, _ := sqlmock.New()
	deleteTx, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	mealItem := &entities.MealItem{
		ID:           id,
		MealID:       mealID,
		MealTypeID:   1,
		ItemName:     "test",
		Qty:          1,
		Unit:         "test",
		Weight:       0,
		Calories:     0,
		Protein:      0,
		Fat:          0,
		Sugars:       0,
		Cholesterol:  0,
		Carbohydrate: 0,
	}

	for idx, tc := range testCases {
		todayMeal := &entities.Meal{
			ID:        tc.todayMealID,
			UserID:    uid,
			MealItems: []entities.MealItem{},
			CalculatedNutrients: entities.CalculatedNutrients{
				TotalCalories:     0,
				TotalCarbohydrate: 0,
				TotalProtein:      0,
				TotalFat:          0,
			},
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/meals/items/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ctx, cancel := context.WithCancel(c.Request().Context())
		defer cancel()

		mockMealRepo := new(mockrepo.MockMealRepository)
		mockMealItemRepo := new(mockrepo.MockMealItemRepository)
		mockMealRepo.On("GetTodayMeal", ctx, uid, start, end).Return(todayMeal, tc.errs[0])
		mockMealItemRepo.On("DeleteMealItem", ctx, id).Return(mealItem, deleteTx, tc.errs[1])

		meal := &entities.Meal{
			ID:     tc.todayMealID,
			UserID: uid,
			CalculatedNutrients: entities.CalculatedNutrients{
				TotalCalories:     0,
				TotalCarbohydrate: 0,
				TotalProtein:      0,
				TotalFat:          0,
				TotalCholesterol:  0,
				TotalSugars:       0,
			},
		}
		mockMealRepo.On("UpdateMeal", ctx, meal).Return(tc.errs[2])

		mealUsecase := usecases.NewMealUsecase(mockMealRepo, mockMealItemRepo)

		err := mealUsecase.DeleteMealItem(c, uid, id)

		if idx == 0 {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}

}

func TestGetUserMeals(t *testing.T) {
	uid := uuid.New()
	p := &dto.PaginationRequest{
		Limit: 1,
		Page:  1,
	}
	meals := &[]entities.Meal{
		{
			ID:     uuid.New(),
			UserID: uid,
		},
	}
	testCases := []testCasePaginationMeal{
		{
			testCase: testCase{
				name: "success",
				errs: []error{nil},
			},
			totalData: 1,
		},
		{
			testCase: testCase{
				name: "get user meal error",
				errs: []error{errors.New("get user meal error")},
			},
			totalData: 1,
		},
		{
			testCase: testCase{
				name: "page overload",
				errs: []error{nil},
			},
			totalData: 0,
		},
	}

	for idx, tc := range testCases {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/meals", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ctx, cancel := context.WithCancel(c.Request().Context())
		defer cancel()

		mockMealRepo := new(mockrepo.MockMealRepository)
		mockMealRepo.On("GetUserMeals", ctx, uid, p).Return(meals, int64(tc.totalData), tc.errs[0])

		mealUsecase := usecases.NewMealUsecase(mockMealRepo, nil)

		_, _, _, err := mealUsecase.GetUserMeals(c, uid, p)

		if idx == 0 {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
