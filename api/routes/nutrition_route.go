package routes

import (
	"nutri-plans-api/api/controllers"
	ex "nutri-plans-api/externals/nutrition"
	"nutri-plans-api/usecases"
	tokenutil "nutri-plans-api/utils/token"
	valutil "nutri-plans-api/utils/validation"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func initNutritionRoute(g *echo.Group, v *valutil.Validator) {
	appID := os.Getenv("NUTRITIONIX_APP_ID")
	appKey := os.Getenv("NUTRITIONIX_APP_KEY")

	nutritionExternal := ex.NewNutritionClient(appKey, appID)
	nutritionUsecase := usecases.NewNutritionUsecase(nutritionExternal)
	nutritionController := controllers.NewNutritionController(nutritionUsecase, v)

	g.Use(echojwt.WithConfig(tokenutil.GetJwtConfig()))

	g.GET("/items/search", nutritionController.SearchItem)
	g.GET("/:item-name", nutritionController.GetItemNutrition)
}
