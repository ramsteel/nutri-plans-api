package dto

type FoodTypeRequest struct {
	Name string `json:"name" validate:"required"`
}
