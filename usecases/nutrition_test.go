package usecases_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"nutri-plans-api/dto"
	exnutri "nutri-plans-api/externals/nutrition"
	mocknutri "nutri-plans-api/mocks/externals/nutrition"
	"nutri-plans-api/usecases"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type testCaseSearchItem struct {
	testCase
	offset int
}

func TestNutritionUsecase(t *testing.T) {
	assert.NotNil(t,
		usecases.NewNutritionUsecase(
			mocknutri.NewMockNutritionClient(t),
		),
	)
}

func TestSearchItem(t *testing.T) {
	items := &[]exnutri.Item{
		{
			ID:   "1",
			Name: "apple",
		},
		{
			ID:   "2",
			Name: "banana",
		},
	}

	testCases := []testCaseSearchItem{
		{
			testCase: testCase{
				name: "success",
				errs: []error{nil},
			},
			offset: 0,
		},
		{
			testCase: testCase{
				name: "offset bigger than total data",
				errs: []error{nil},
			},
			offset: 3,
		},
		{
			testCase: testCase{
				name: "sum offset and limit bigger than total data",
				errs: []error{nil},
			},
			offset: 2,
		},
		{
			testCase: testCase{
				name: "search item error",
				errs: []error{errors.New("search item error")},
			},
			offset: 0,
		},
	}

	for idx, tc := range testCases {
		s := &dto.SearchRequest{
			Item:   "apple",
			Offset: &tc.offset,
			Limit:  1,
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/nutrition/items", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ctx, cancel := context.WithCancel(c.Request().Context())
		defer cancel()

		mockNutritionClient := new(mocknutri.MockNutritionClient)
		mockNutritionClient.On("SearchItem", ctx, s.Item).Return(items, tc.errs[0])

		nutritionUsecase := usecases.NewNutritionUsecase(mockNutritionClient)

		_, _, err := nutritionUsecase.SearchItem(c, s)
		if idx == len(testCases)-1 {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}

}

func TestGetNutritionItem(t *testing.T) {
	r := &dto.ItemNutritionRequest{
		Query: "apple",
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/nutrition/items/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	mockNutritionClient := new(mocknutri.MockNutritionClient)
	mockNutritionClient.On("GetItemNutrition", ctx, r).Return(&exnutri.ItemNutrition{}, nil)

	nutritionUsecase := usecases.NewNutritionUsecase(mockNutritionClient)

	_, err := nutritionUsecase.GetItemNutrition(c, r)
	assert.NoError(t, err)
}
