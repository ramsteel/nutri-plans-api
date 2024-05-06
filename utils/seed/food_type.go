package seed

import "nutri-plans-api/entities"

func GetFoodTypes() *[]entities.FoodType {
	foodTypes := &[]entities.FoodType{
		{
			ID:   1,
			Name: "dry",
		},
		{
			ID:   2,
			Name: "soupy",
		},
		{
			ID:   3,
			Name: "fried",
		},
		{
			ID:   4,
			Name: "starchy",
		},
		{
			ID:   5,
			Name: "grilled",
		},
		{
			ID:   6,
			Name: "baked",
		},
		{
			ID:   7,
			Name: "creamy",
		},
	}
	return foodTypes
}
