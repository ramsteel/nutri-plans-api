package usecases

import (
	"context"
	"nutri-plans-api/dto"
	ex "nutri-plans-api/externals/nutrition"

	"github.com/labstack/echo/v4"
)

type NutritionUsecase interface {
	SearchItem(c echo.Context, itemName string, limit, offset int) (
		*[]ex.Item, *dto.MetadataResponse, error)
	GetItemNutrition(c echo.Context, r *dto.ItemNutritionRequest) (
		*ex.ItemNutrition, error)
}

type nutritionUsecase struct {
	nutrtionExternal ex.NutritionClient
}

func NewNutritionUsecase(nutrtionExternal ex.NutritionClient) *nutritionUsecase {
	return &nutritionUsecase{
		nutrtionExternal: nutrtionExternal,
	}
}

func (n *nutritionUsecase) SearchItem(
	c echo.Context,
	itemName string,
	limit, offset int,
) (*[]ex.Item, *dto.MetadataResponse, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	items, err := n.nutrtionExternal.SearchItem(ctx, itemName)
	if err != nil {
		return nil, nil, err
	}

	items = n.filterItems(items)
	totalData := len(*items)
	if offset > totalData {
		defaultMetadata := &dto.MetadataResponse{
			TotalData:   totalData,
			TotalCount:  totalData,
			NextOffset:  totalData,
			HasLoadMore: totalData > offset+limit,
		}
		return &[]ex.Item{}, defaultMetadata, nil
	}

	if offset+limit > totalData {
		limit = totalData - offset
	}

	filteredItems := (*items)[offset : offset+limit]
	metadata := &dto.MetadataResponse{
		TotalData:   totalData,
		TotalCount:  offset + len(filteredItems),
		NextOffset:  offset + limit,
		HasLoadMore: totalData > offset+limit,
	}

	return &filteredItems, metadata, nil
}

func (n *nutritionUsecase) GetItemNutrition(
	c echo.Context,
	r *dto.ItemNutritionRequest,
) (*ex.ItemNutrition, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return n.nutrtionExternal.GetItemNutrition(ctx, r)
}

func (n *nutritionUsecase) filterItems(items *[]ex.Item) *[]ex.Item {
	idFlags := make(map[string]bool)
	var filteredItems []ex.Item

	for _, item := range *items {
		if _, ok := idFlags[item.ID]; !ok {
			idFlags[item.ID] = true
			filteredItems = append(filteredItems, item)
		}
	}

	return &filteredItems
}
