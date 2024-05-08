package routes

import (
	"nutri-plans-api/api/controllers"
	"nutri-plans-api/repositories"
	"nutri-plans-api/usecases"
	tokenutil "nutri-plans-api/utils/token"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initMealRoute(g *echo.Group, db *gorm.DB) {
	mealRepository := repositories.NewMealRepository(db)
	mealUsecase := usecases.NewMealUsecase(mealRepository)
	tokenUtil := tokenutil.NewTokenUtil()

	mealController := controllers.NewMealController(mealUsecase, tokenUtil)

	g.Use(echojwt.WithConfig(tokenutil.GetJwtConfig()))

	g.GET("/today", mealController.GetTodayMeal)
}
