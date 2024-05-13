package routes

import (
	"nutri-plans-api/api/controllers"
	"nutri-plans-api/repositories"
	"nutri-plans-api/usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initDietaryPreferenceTypeRoute(g *echo.Group, db *gorm.DB) {
	dietaryPreferenceTypeRepository := repositories.NewDietaryPreferenceTypeRepository(db)
	dietaryPreferenceTypeUsecase := usecases.NewDietaryPreferenceTypeUsecase(
		dietaryPreferenceTypeRepository,
	)
	dietaryPreferenceTypeController := controllers.NewDietaryPreferenceTypeController(
		dietaryPreferenceTypeUsecase,
	)

	g.GET("", dietaryPreferenceTypeController.GetDietaryPreferenceTypes)
}
