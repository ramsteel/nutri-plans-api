package routes

import (
	"nutri-plans-api/api/controllers"
	"nutri-plans-api/repositories"
	"nutri-plans-api/usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initCountryRoute(g *echo.Group, db *gorm.DB) {
	countryRepository := repositories.NewCountryRepository(db)
	countryUsecase := usecases.NewCountryUsecase(countryRepository)
	countryController := controllers.NewCountryController(countryUsecase)

	g.GET("", countryController.GetCountries)
}
