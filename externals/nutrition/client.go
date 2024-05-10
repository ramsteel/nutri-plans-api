package nutrition

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"nutri-plans-api/dto"
	errutil "nutri-plans-api/utils/error"
)

type NutritionClient interface {
	SearchItem(ctx context.Context, name string) (*[]Item, error)
	GetItemNutrition(ctx context.Context, r *dto.ItemNutritionRequest) (*ItemNutrition, error)
	GetMultipleItemNutritions(
		ctx context.Context,
		r *dto.ItemNutritionRequest,
	) (*[]ItemNutrition, error)
}

type nutritionClient struct {
	appKey string
	appID  string
	client *http.Client
}

func NewNutritionClient(appKey, appID string) *nutritionClient {
	return &nutritionClient{
		appKey: appKey,
		appID:  appID,
		client: http.DefaultClient,
	}
}

func (n *nutritionClient) SearchItem(ctx context.Context, name string) (*[]Item, error) {
	url := "https://trackapi.nutritionix.com/v2/search/instant?branded=false&query=" + name

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errutil.ErrExternalService
	}

	req.Header.Set("x-app-id", n.appID)
	req.Header.Set("x-app-key", n.appKey)
	res, err := n.client.Do(req)
	if err != nil {
		return nil, errutil.ErrExternalService
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, errutil.ErrItemNotFound
	}

	searchItemRes := new(SearchItemResponse)
	err = json.NewDecoder(res.Body).Decode(searchItemRes)
	if err != nil {
		return nil, err
	}

	return searchItemRes.Common, nil
}

func (n *nutritionClient) GetItemNutrition(
	ctx context.Context,
	r *dto.ItemNutritionRequest,
) (*ItemNutrition, error) {
	nutritionRes, err := n.getItemNutrition(ctx, r)
	if err != nil {
		return nil, err
	}

	return &(*nutritionRes.Foods)[0], nil
}

func (n *nutritionClient) getItemNutrition(
	ctx context.Context,
	r *dto.ItemNutritionRequest,
) (*NutritionResponse, error) {
	url := "https://trackapi.nutritionix.com/v2/natural/nutrients"

	strReq, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		strings.NewReader(string(strReq)),
	)
	if err != nil {
		return nil, errutil.ErrExternalService
	}

	req.Header.Set("x-app-id", n.appID)
	req.Header.Set("x-app-key", n.appKey)

	res, err := n.client.Do(req)
	if err != nil {
		return nil, errutil.ErrExternalService
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errutil.ErrItemNotFound
	}

	nutritionRes := new(NutritionResponse)
	err = json.NewDecoder(res.Body).Decode(nutritionRes)
	if err != nil {
		return nil, err
	}

	return nutritionRes, nil
}

func (n *nutritionClient) GetMultipleItemNutritions(
	ctx context.Context,
	r *dto.ItemNutritionRequest,
) (*[]ItemNutrition, error) {
	nutritionRes, err := n.getItemNutrition(ctx, r)
	if err != nil {
		return nil, err
	}

	return (*nutritionRes).Foods, nil
}
