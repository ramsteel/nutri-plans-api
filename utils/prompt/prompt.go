package prompt

import (
	"fmt"
	"nutri-plans-api/entities"
)

func GetRecommendationPrompt(p *entities.UserPreference) string {
	start := "You are food and drink recommender and give seven list of healthy food/drink/both randomly based on this preference:\n"

	pref := "- Food type: %s\n- Drink type: %s\n- Dietary preference: %s\n- Dietary restriction: %s\n"
	var (
		foodType         = "any"
		drinkType        = "any"
		dietType         = "any"
		dietRestrictions = "-"
	)

	if p.FoodType != nil {
		foodType = p.FoodType.Name
	}

	if p.DrinkType != nil {
		drinkType = p.DrinkType.Name
	}

	if p.DietaryPreferenceType != nil {
		dietType = p.DietaryPreferenceType.Name
	}

	if len(*p.DietaryRestrictions) > 0 {
		restrictions := ""
		for _, v := range *p.DietaryRestrictions {
			restrictions += fmt.Sprint(v.Name, ", ")
		}

		dietRestrictions = restrictions[:len(restrictions)-2]
	}

	pref = fmt.Sprintf(pref, foodType, drinkType, dietType, dietRestrictions)

	end := "\n*Response specification : only name without any words at the beginning.*"

	prompt := fmt.Sprint(start, pref, end)

	return prompt
}
