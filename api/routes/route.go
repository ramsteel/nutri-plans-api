package routes

import (
	valutil "nutri-plans-api/utils/validation"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB, v *valutil.Validator) {
	baseRoute := e.Group("/api/v1")

	userRoute := baseRoute.Group("")
	userPreference := baseRoute.Group("/preference")
	countryRoute := baseRoute.Group("/countries")
	foodTypeRoute := baseRoute.Group("/food-types")
	drinkTypeRoute := baseRoute.Group("/drink-types")
	dietaryPreferenceTypeRoute := baseRoute.Group("/dietary-preference-types")
	nutritionRoute := baseRoute.Group("/nutrition")
	mealTypeRoute := baseRoute.Group("/meal-types")
	mealRoute := baseRoute.Group("/meals")
	recommendationRoute := baseRoute.Group("/recommendation")
	adminRoute := baseRoute.Group("/admin")

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
