package seed

import "nutri-plans-api/entities"

func GetDietaryPreferenceTypes() *[]entities.DietaryPreferenceType {
	dietaryPreferenceTypes := &[]entities.DietaryPreferenceType{
		{
			Name:        "vegan",
			Description: "vegan is a type of food that does not contain any animal products. ex: beef, chicken, tofu, etc.",
		},
		{
			Name:        "vegetarian",
			Description: "vegetarian is a type of food that does not contain any animal products and that does not contain any egg products. ex: broccoli, kale, spinach, etc.",
		},
		{
			Name:        "gluten-free",
			Description: "gluten-free is a type of food that does not contain any gluten products. ex: bread, cereal, pasta, etc.",
		},
		{
			Name:        "dairy-free",
			Description: "dairy-free is a type of food that does not contain any dairy products. Ex: milk, cheese, yogurt, etc.",
		},
	}
	return dietaryPreferenceTypes
}
