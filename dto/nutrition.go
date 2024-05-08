package dto

type ItemNutritionRequest struct {
	Query string `json:"query" validate:"required"`
}
