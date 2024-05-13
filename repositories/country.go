package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type CountryRepository interface {
	GetAllCountries(ctx context.Context) (*[]entities.Country, error)
	GetCountryByID(ctx context.Context, id uint) (*entities.Country, error)
}

type countryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *countryRepository {
	return &countryRepository{
		db: db,
	}
}

func (r *countryRepository) GetCountryByID(ctx context.Context, id uint) (*entities.Country, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	country := new(entities.Country)
	if err := r.db.First(country, id).Error; err != nil {
		return nil, err
	}

	return country, nil
}

func (r *countryRepository) GetAllCountries(ctx context.Context) (*[]entities.Country, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	countries := new([]entities.Country)
	if err := r.db.Find(countries).Error; err != nil {
		return nil, err
	}

	return countries, nil
}
