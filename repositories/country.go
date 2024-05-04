package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type CountryRepository interface {
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
