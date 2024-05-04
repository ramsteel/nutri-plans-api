package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateAuth(ctx context.Context, auth *entities.Auth) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (a *authRepository) CreateAuth(ctx context.Context, auth *entities.Auth) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	return a.db.Create(auth).Error
}
