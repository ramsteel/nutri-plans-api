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
	adminUsecase := usecases.NewAdminUsecase(adminRepository)
	tokenUtil := tokenutil.NewTokenUtil()

	adminContoller := controllers.NewAdminController(adminUsecase, tokenUtil, v)

	g.Use(echojwt.WithConfig(tokenutil.GetJwtConfig()), middlewares.MustAdmin)

	g.GET("", adminContoller.GetAdminProfile)
}
