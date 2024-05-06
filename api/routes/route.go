package routes

import (
	valutil "nutri-plans-api/utils/validation"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB, v *valutil.Validator) {
	userRoute := e.Group("")
	userPreference := e.Group("/preferences")

	initUserRoute(userRoute, db, v)
	initUserPreferenceRoute(userPreference, db, v)
}
