package dto

type UserPreferenceRequest struct {
	FoodTypeID              *uint     `json:"food_type_id"`
	DrinkTypeID             *uint     `json:"drink_type_id"`
	DietaryPreferenceTypeID *uint     `json:"dietary_preference_type_id"`
	DietaryRestrictions     *[]string `json:"dietary_restrictions"`
}
