package seed

import "nutri-plans-api/entities"

func GetDrinkTypes() *[]entities.DrinkType {
	drinkTypes := &[]entities.DrinkType{
		{
			Name: "hot",
		},
		{
			Name: "cold",
		},
		{
			Name: "juice",
		},
		{
			Name: "smoothie",
		},
		{
			Name: "milk-based",
		},
		{
			Name: "infused water",
		},
	}
	return drinkTypes
}
