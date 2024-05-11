package dto

type DrinkTypeRequest struct {
	Name string `json:"name" validate:"required"`
}
