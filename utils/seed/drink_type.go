package seed

import "nutri-plans-api/entities"

func GetDrinkTypes() *[]entities.DrinkType {
	drinkTypes := &[]entities.DrinkType{
		{
			ID:   1,
			Name: "hot",
		},
		{
			ID:   2,
			Name: "cold",
		},
		{
			ID:   3,
			Name: "juice",
		},
		{
			ID:   4,
			Name: "smoothie",
		},
		{
			ID:   5,
			Name: "milk-based",
		},
		{
			ID:   6,
			Name: "infused water",
		},
	}
	return drinkTypes
}
