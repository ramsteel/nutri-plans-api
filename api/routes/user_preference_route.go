package routes

import (
	"nutri-plans-api/api/controllers"
	"nutri-plans-api/repositories"
	"nutri-plans-api/usecases"
	tokenutil "nutri-plans-api/utils/token"
	valutil "nutri-plans-api/utils/validation"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initUserPreferenceRoute(g *echo.Group, db *gorm.DB, v *valutil.Validator) {
	userPreferenceRepository := repositories.NewUserPreferenceRepository(db)
	dietaryRestRepository := repositories.NewDietaryRestrictionRepository(db)
	userPreferenceUsecase := usecases.NewUserPreferenceUsecase(
		userPreferenceRepository,
		dietaryRestRepository,
	)
	tokenUtil := tokenutil.NewTokenUtil()

	userPreferenceController := controllers.NewUserPreferenceController(
		userPreferenceUsecase,
		v,
		tokenUtil,
	)

	g.Use(echojwt.WithConfig(tokenutil.GetJwtConfig()))

	g.PUT("", userPreferenceController.UpdateUserPreference)
	g.GET("", userPreferenceController.GetUserPreference)
}
