package routes

import (
	"nutri-plans-api/api/controllers"
	"nutri-plans-api/repositories"
	"nutri-plans-api/usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initFoodTypeRoute(g *echo.Group, db *gorm.DB) {
	foodTypeRepository := repositories.NewFoodTypeRepository(db)
	foodTypeUsecase := usecases.NewFoodTypeUsecase(foodTypeRepository)
	foodTypeController := controllers.NewFoodTypeController(foodTypeUsecase)

	g.GET("", foodTypeController.GetFoodTypes)
}
