package usecases_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"nutri-plans-api/dto"
	"nutri-plans-api/entities"
	exnutri "nutri-plans-api/externals/nutrition"
	mocknutri "nutri-plans-api/mocks/externals/nutrition"
	mockoai "nutri-plans-api/mocks/externals/openai"
	mockrepo "nutri-plans-api/mocks/repositories"
	"nutri-plans-api/usecases"
	"nutri-plans-api/utils/prompt"
	recutil "nutri-plans-api/utils/recommendation"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewRecommendationUsecase(t *testing.T) {
	assert.NotNil(
		t,
		usecases.NewRecommendationUsecase(
			mockrepo.NewMockRecommendationRepository(t),
			mockrepo.NewMockUserPreferenceRepository(t),
			mocknutri.NewMockNutritionClient(t),
			mockoai.NewMockOpenAIClient(t),
		),
	)
}

func TestGetRecommendation(t *testing.T) {
	uid := uuid.New()
	recommendations := &[]entities.Recommendation{
		{
			ID:               1,
			UserPreferenceID: uid,
			Name:             "apple",
		},
	}
	itemReq := &dto.ItemNutritionRequest{
		Query: "apple",
	}
	itemNutritions := &[]exnutri.ItemNutrition{
		{
			ItemName:      "apple",
			ServingQty:    1,
			ServingUnit:   "serving",
			ServingWeight: 123,
			Nutrient: exnutri.Nutrient{
				Calories:     1,
				Fat:          1,
				Protein:      1,
				Carbohydrate: 1,
				Cholesterol:  1,
				Sugar:        1,
			},
		},
	}

	testCases := []testCase{
		{
			name: "success",
			errs: []error{nil, nil},
		},
		{
			name: "get recommendation error",
			errs: []error{errors.New("get recommendation error"), nil},
		},
		{
			name: "get multiple item nutrition error",
			errs: []error{nil, errors.New("get multiple item nutrition error")},
		},
	}

	for idx, tc := range testCases {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/recommendation", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ctx, cancel := context.WithCancel(c.Request().Context())
		defer cancel()

		mockRecRepo := new(mockrepo.MockRecommendationRepository)
		mockRecRepo.On("GetRecommendations", ctx, uid).Return(recommendations, tc.errs[0])

		mockNutritionClient := new(mocknutri.MockNutritionClient)
		mockNutritionClient.On(
			"GetMultipleItemNutritions",
			ctx,
			itemReq,
		).Return(itemNutritions, tc.errs[1])

		recommendationUsecase := usecases.NewRecommendationUsecase(
			mockRecRepo,
			nil,
			mockNutritionClient,
			nil,
		)

		_, err := recommendationUsecase.GetRecommendation(c, uid)
		if idx == 0 {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}

}

func TestFetchOpenAIRecommendation(t *testing.T) {
	prefID := uuid.New()
	preference := &[]entities.UserPreference{
		{
			UserID: prefID,
			FoodType: &entities.FoodType{
				ID:   1,
				Name: "food type",
			},
			DrinkType: &entities.DrinkType{
				ID:   1,
				Name: "drink type",
			},
			DietaryPreferenceType: &entities.DietaryPreferenceType{
				ID:          1,
				Name:        "dietary preference type",
				Description: "description",
			},
			DietaryRestrictions: &[]entities.DietaryRestriction{},
			Recommendations: &[]entities.Recommendation{
				{
					ID:               1,
					UserPreferenceID: prefID,
					Name:             "apple",
				},
			},
		},
	}

	mockUserPrefRepo := new(mockrepo.MockUserPreferenceRepository)
	mockUserPrefRepo.On("GetAllUserPreferences", context.Background()).Return(preference, nil)

	testCases := []testCase{
		{
			name: "success",
			errs: []error{nil},
		},
		{
			name: "create recommendation error",
			errs: []error{errors.New("create recommendation error")},
		},
	}

	for _, pref := range *preference {
		for _, tc := range testCases {
			mockOaiClient := new(mockoai.MockOpenAIClient)
			exPrompt := prompt.GetRecommendationPrompt(&pref)
			res := "1. apple"
			mockOaiClient.On(
				"GetRecommendation",
				exPrompt,
				recutil.ToString(pref.Recommendations, true),
			).Return(res, nil)

			mockRecRepo := new(mockrepo.MockRecommendationRepository)
			recommendations := recutil.ToStruct(res, pref.UserID)
			mockRecRepo.On(
				"CreateRecommendations",
				context.Background(),
				recommendations,
			).Return(tc.errs[0])

			recommendationUsecase := usecases.NewRecommendationUsecase(
				mockRecRepo,
				mockUserPrefRepo,
				nil,
				mockOaiClient,
			)

			recommendationUsecase.FetchOpenAIRecommendation()
		}
	}
}
