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

func initMealRoute(g *echo.Group, db *gorm.DB, v *valutil.Validator) {
	mealRepository := repositories.NewMealRepository(db)
	mealItemRepository := repositories.NewMealItemRepository(db)
	mealUsecase := usecases.NewMealUsecase(mealRepository, mealItemRepository)
	tokenUtil := tokenutil.NewTokenUtil()

	mealController := controllers.NewMealController(mealUsecase, tokenUtil, v)

	g.Use(echojwt.WithConfig(tokenutil.GetJwtConfig()))

	g.GET("/today", mealController.GetTodayMeal)
	g.POST("/items", mealController.AddItemToMeal)
	g.PUT("/items/:id", mealController.UpdateMealItem)
	g.GET("/items/:id", mealController.GetMealItemByID)
}
