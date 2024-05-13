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
	drinkTypeRepository := repositories.NewDrinkTypeRepository(db)
	dietaryPreferenceTypeRepository := repositories.NewDietaryPreferenceTypeRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	adminUsecase := usecases.NewAdminUsecase(adminRepository)
	foodTypeUsecase := usecases.NewFoodTypeUsecase(foodTypeRepository)
	drinkTypeUsecase := usecases.NewDrinkTypeUsecase(drinkTypeRepository)
	dietaryPreferenceTypeUsecase := usecases.NewDietaryPreferenceTypeUsecase(
		dietaryPreferenceTypeRepository,
	)
	authUsecase := usecases.NewAuthUsecase(authRepository)

	tokenUtil := tokenutil.NewTokenUtil()

	adminContoller := controllers.NewAdminController(
		adminUsecase,
		foodTypeUsecase,
		drinkTypeUsecase,
		dietaryPreferenceTypeUsecase,
		authUsecase,
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

	// drink types
	g.POST("/drink-types", adminContoller.CreateDrinkType)
	g.PUT("/drink-types/:id", adminContoller.UpdateDrinkType)
	g.DELETE("/drink-types/:id", adminContoller.DeleteDrinkType)

	// dietary preference type
	g.POST("/dietary-preference-types", adminContoller.CreateDietaryPreferenceType)
	g.PUT("/dietary-preference-types/:id", adminContoller.UpdateDietaryPreferenceType)
	g.DELETE("/dietary-preference-types/:id", adminContoller.DeleteDietaryPreferenceType)

	// users
	g.GET("/users", adminContoller.GetAllUsersAuth)
}
