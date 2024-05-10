package seed

import "nutri-plans-api/entities"

func GetMealTypes() *[]entities.MealType {
	mealTypes := &[]entities.MealType{
		{
			ID:   1,
			Name: "breakfast",
		},
		{
			ID:   2,
			Name: "lunch",
		},
		{
			ID:   3,
			Name: "dinner",
		},
		{
			ID:   4,
			Name: "additional",
		},
	}
	return mealTypes
}
