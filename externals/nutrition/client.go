package nutrition

import (
	"context"
	"encoding/json"
	"net/http"

	errutil "nutri-plans-api/utils/error"
)

type NutritionClient interface {
	SearchItem(ctx context.Context, name string) (*[]Item, error)
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
		return nil, errutil.ErrFailedDecodeJson
	}

	return searchItemRes.Common, nil
}
