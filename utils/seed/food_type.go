package seed

import "nutri-plans-api/entities"

func GetFoodTypes() *[]entities.FoodType {
	foodTypes := &[]entities.FoodType{
		{
			Name: "dry",
		},
		{
			Name: "soupy",
		},
		{
			Name: "fried",
		},
		{
			Name: "starchy",
		},
		{
			Name: "grilled",
		},
		{
			Name: "baked",
		},
		{
			Name: "creamy",
		},
	}
	return foodTypes
}
