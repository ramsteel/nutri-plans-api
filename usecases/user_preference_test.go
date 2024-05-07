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
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type filterDietaryRestTestCase struct {
	name                string
	dietaryRestrictions *[]entities.DietaryRestriction
	data                *[]string
}

func TestNewUserPreferenceUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewUserPreferenceUsecase(
			mockrepo.NewMockUserPreferenceRepository(t),
			mockrepo.NewMockDietaryRestrictionRepository(t),
		),
	)
}

func TestUpdateUserPreference(t *testing.T) {
	temp := uint(1)
	var typeID *uint = &temp
	var (
		id = uuid.New()
		r  = &dto.UserPreferenceRequest{
			FoodTypeID:              typeID,
			DrinkTypeID:             typeID,
			DietaryPreferenceTypeID: typeID,
			DietaryRestrictions:     &[]string{"egg"},
		}
		dietaryRestrictions = &[]entities.DietaryRestriction{
			{
				ID:               1,
				UserPreferenceID: id,
				Name:             "egg",
				CreatedAt:        time.UnixMilli(1714757476909),
				UpdatedAt:        time.UnixMilli(1714757476909),
				DeletedAt:        gorm.DeletedAt{},
			},
		}

		userPreference = &entities.UserPreference{
			UserID:                  id,
			FoodTypeID:              r.FoodTypeID,
			DrinkTypeID:             r.DrinkTypeID,
			DietaryPreferenceTypeID: r.DietaryPreferenceTypeID,
			DietaryRestrictions:     dietaryRestrictions,
		}
	)

	testCases := []testCase{
		{
			name: "success",
			errs: []error{nil, nil},
		},
		{
			name: "error delete dietary restrictions",
			errs: []error{errors.New("failed delete"), nil},
		},
		{
			name: "error get dietary restrictions",
			errs: []error{nil, errors.New("failed get")},
		},
	}

	for idx, tc := range testCases {
		mockDietaryRestrictionRepo := new(mockrepo.MockDietaryRestrictionRepository)
		mockUserPreferenceRepo := new(mockrepo.MockUserPreferenceRepository)
		u := usecases.NewUserPreferenceUsecase(
			mockUserPreferenceRepo,
			mockDietaryRestrictionRepo,
		)
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/preferences", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctx, cancel := context.WithCancel(c.Request().Context())
			defer cancel()
			mockDietaryRestrictionRepo.On("DeleteDietaryRestriction", ctx, id).Return(
				tc.errs[0],
			)
			mockDietaryRestrictionRepo.On("GetDeletedDietaryRestrictions", ctx, id).Return(
				dietaryRestrictions,
				tc.errs[1],
			)
			mockUserPreferenceRepo.On("UpdateUserPreference", ctx, userPreference).Return(nil)
			err := u.UpdateUserPreference(c, id, r)

			if idx != 0 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetUserPreference(t *testing.T) {
	var (
		uid            = uuid.New()
		userPreference = &entities.UserPreference{
			UserID:                  uid,
			FoodTypeID:              nil,
			DrinkTypeID:             nil,
			DietaryPreferenceTypeID: nil,
			DietaryRestrictions:     nil,
		}
	)
	mockGetUserPreferenceRepo := new(mockrepo.MockUserPreferenceRepository)
	u := usecases.NewUserPreferenceUsecase(
		mockGetUserPreferenceRepo,
		nil,
	)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/preferences", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	mockGetUserPreferenceRepo.On("GetUserPreference", ctx, uid).Return(userPreference, nil)
	defer cancel()
	res, err := u.GetUserPreference(c, uid)
	assert.NoError(t, err)
	assert.Equal(t, res, userPreference)
}

func TestFilterDietaryRestrictions(t *testing.T) {
	var (
		uid       = uuid.New()
		testCases = []filterDietaryRestTestCase{
			{
				name: "same",
				dietaryRestrictions: &[]entities.DietaryRestriction{
					{
						ID:               1,
						UserPreferenceID: uid,
						Name:             "egg",
						CreatedAt:        time.UnixMilli(1714757476909),
						UpdatedAt:        time.UnixMilli(1714757476909),
						DeletedAt:        gorm.DeletedAt{},
					},
				},
				data: &[]string{"egg"},
			},
			{
				name: "same but with duplicate",
				dietaryRestrictions: &[]entities.DietaryRestriction{
					{
						ID:               1,
						UserPreferenceID: uid,
						Name:             "egg",
						CreatedAt:        time.UnixMilli(1714757476909),
						UpdatedAt:        time.UnixMilli(1714757476909),
						DeletedAt:        gorm.DeletedAt{},
					},
				},
				data: &[]string{"egg", "egg"},
			},
			{
				name: "different",
				dietaryRestrictions: &[]entities.DietaryRestriction{
					{
						ID:               1,
						UserPreferenceID: uid,
						Name:             "egg",
						CreatedAt:        time.UnixMilli(1714757476909),
						UpdatedAt:        time.UnixMilli(1714757476909),
						DeletedAt:        gorm.DeletedAt{},
					},
				},
				data: &[]string{"egg", "beef"},
			},
		}
	)
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			usecases.FilterDietaryRestrictions(tt.dietaryRestrictions, tt.data)
		})
	}
}
