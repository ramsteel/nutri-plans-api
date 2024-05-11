package routes

import (
	valutil "nutri-plans-api/utils/validation"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB, v *valutil.Validator) {
	userRoute := e.Group("")
	userPreference := e.Group("/preference")
	countryRoute := e.Group("/countries")
	foodTypeRoute := e.Group("/food-types")
	drinkTypeRoute := e.Group("/drink-types")
	dietaryPreferenceTypeRoute := e.Group("/dietary-preference-types")
	nutritionRoute := e.Group("/nutrition")
	mealTypeRoute := e.Group("/meal-types")
	mealRoute := e.Group("/meals")
	recommendationRoute := e.Group("/recommendation")
	adminRoute := e.Group("/admin")

	initUserRoute(userRoute, db, v)
	initUserPreferenceRoute(userPreference, db, v)
	initCountryRoute(countryRoute, db)
	initFoodTypeRoute(foodTypeRoute, db)
	initDrinkTypeRoute(drinkTypeRoute, db)
	initDietaryPreferenceTypeRoute(dietaryPreferenceTypeRoute, db)
	initNutritionRoute(nutritionRoute, v)
	initMealTypeRoute(mealTypeRoute, db)
	initMealRoute(mealRoute, db, v)
	initRecommendationRoute(recommendationRoute, db)
	initAdminRoute(adminRoute, db, v)
}
