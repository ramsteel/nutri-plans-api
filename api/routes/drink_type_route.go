package routes

import (
	"nutri-plans-api/api/controllers"
	"nutri-plans-api/repositories"
	"nutri-plans-api/usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initDrinkTypeRoute(g *echo.Group, db *gorm.DB) {
	drinkTypeRepository := repositories.NewDrinkTypeRepository(db)
	drinkTypeUsecase := usecases.NewDrinkTypeUsecase(drinkTypeRepository)
	drinkTypeController := controllers.NewDrinkTypeController(drinkTypeUsecase)

	g.GET("", drinkTypeController.GetDrinkTypes)
}
