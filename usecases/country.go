package usecases

import (
	"context"
	"nutri-plans-api/entities"
	"nutri-plans-api/repositories"

	"github.com/labstack/echo/v4"
)

type CountryUsecase interface {
	GetCountries(c echo.Context) (*[]entities.Country, error)
}

type countryUsecase struct {
	countryRepo repositories.CountryRepository
}

func NewCountryUsecase(countryRepo repositories.CountryRepository) *countryUsecase {
	return &countryUsecase{
		countryRepo: countryRepo,
	}
}

func (u *countryUsecase) GetCountries(c echo.Context) (*[]entities.Country, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	return u.countryRepo.GetAllCountries(ctx)
}
