package routes

import (
	"nutri-plans-api/api/controllers"
	"nutri-plans-api/api/middlewares"
	"nutri-plans-api/repositories"
	"nutri-plans-api/usecases"
	tokenutil "nutri-plans-api/utils/token"
	valutil "nutri-plans-api/utils/validation"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initAdminRoute(g *echo.Group, db *gorm.DB, v *valutil.Validator) {
	adminRepository := repositories.NewAdminRepository(db)
	foodTypeRepository := repositories.NewFoodTypeRepository(db)

	adminUsecase := usecases.NewAdminUsecase(adminRepository)
	foodTypeUsecase := usecases.NewFoodTypeUsecase(foodTypeRepository)
	tokenUtil := tokenutil.NewTokenUtil()

	adminContoller := controllers.NewAdminController(
		adminUsecase,
		foodTypeUsecase,
		tokenUtil,
		v,
	)

	g.Use(echojwt.WithConfig(tokenutil.GetJwtConfig()), middlewares.MustAdmin)

	// profile
	g.GET("", adminContoller.GetAdminProfile)

	// food types
	g.POST("/food-types", adminContoller.CreateFoodType)
	g.PUT("/food-types/:id", adminContoller.UpdateFoodType)
	g.DELETE("/food-types/:id", adminContoller.DeleteFoodType)
}
