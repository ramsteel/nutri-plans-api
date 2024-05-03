package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(ctx context.Context, user *entities.User) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	return u.db.Create(user).Error
}
