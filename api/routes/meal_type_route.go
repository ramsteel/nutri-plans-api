package routes

import (
	"nutri-plans-api/api/controllers"
	"nutri-plans-api/repositories"
	"nutri-plans-api/usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initMealTypeRoute(g *echo.Group, db *gorm.DB) {
	mealTypeRepository := repositories.NewMealTypeRepository(db)
	mealTypeUsecase := usecases.NewMealTypeUsecase(mealTypeRepository)
	mealTypeController := controllers.NewMealTypeController(mealTypeUsecase)

	g.GET("", mealTypeController.GetMealTypes)
}
