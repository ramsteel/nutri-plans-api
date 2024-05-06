package routes

import (
	valutil "nutri-plans-api/utils/validation"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB, v *valutil.Validator) {
	userRoute := e.Group("")
	userPreference := e.Group("/preferences")
	countryRoute := e.Group("/countries")
	foodTypeRoute := e.Group("/food-types")
	drinkTypeRoute := e.Group("/drink-types")

	initUserRoute(userRoute, db, v)
	initUserPreferenceRoute(userPreference, db, v)
	initCountryRoute(countryRoute, db)
	initFoodTypeRoute(foodTypeRoute, db)
	initDrinkTypeRoute(drinkTypeRoute, db)
}
