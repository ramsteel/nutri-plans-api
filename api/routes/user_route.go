package routes

import (
	"nutri-plans-api/api/controllers"
	"nutri-plans-api/repositories"
	"nutri-plans-api/usecases"
	passutil "nutri-plans-api/utils/password"
	tokenutil "nutri-plans-api/utils/token"
	valutil "nutri-plans-api/utils/validation"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initUserRoute(g *echo.Group, db *gorm.DB, v *valutil.Validator) {
	userRepository := repositories.NewUserRepository(db)
	countryRepository := repositories.NewCountryRepository(db)
	authRepository := repositories.NewAuthRepository(db)
	passUtil := passutil.NewPasswordUtil()
	tokenUtil := tokenutil.NewTokenUtil()

	userUsecase := usecases.NewUserUsecase(
		userRepository,
		authRepository,
		countryRepository,
		passUtil,
		tokenUtil,
	)
	userController := controllers.NewUserController(userUsecase, v, tokenUtil)

	g.POST("/register", userController.Register)
	g.POST("/login", userController.Login)

	g.Use(echojwt.WithConfig(tokenutil.GetJwtConfig()))
	g.GET("/profiles", userController.GetUserDetail)
	g.PUT("/profiles", userController.UpdateUser)
}
