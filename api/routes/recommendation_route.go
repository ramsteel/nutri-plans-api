package routes

import (
	"nutri-plans-api/api/controllers"
	exnutri "nutri-plans-api/externals/nutrition"
	exoai "nutri-plans-api/externals/openai"
	"nutri-plans-api/repositories"
	"nutri-plans-api/usecases"
	tokenutil "nutri-plans-api/utils/token"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initRecommendationRoute(g *echo.Group, db *gorm.DB) {
	nixID := os.Getenv("NUTRITIONIX_APP_ID")
	nixKey := os.Getenv("NUTRITIONIX_APP_KEY")
	oaiKey := os.Getenv("OPENAI_API_KEY")

	recommendationRepository := repositories.NewRecommendationRepository(db)
	userPreferenceRepository := repositories.NewUserPreferenceRepository(db)
	nutritionExternal := exnutri.NewNutritionClient(nixKey, nixID)
	oaiExternal := exoai.NewOpenAIClient(oaiKey)

	tokenUtil := tokenutil.NewTokenUtil()

	recommendationUsecase := usecases.NewRecommendationUsecase(
		recommendationRepository,
		userPreferenceRepository,
		nutritionExternal,
		oaiExternal,
	)

	recommendationUsecase.StartRecommendationCron()

	recommendationController := controllers.NewRecommendationController(
		recommendationUsecase,
		tokenUtil,
	)

	g.Use(echojwt.WithConfig(tokenutil.GetJwtConfig()))

	g.GET("", recommendationController.GetRecommendation)
}
