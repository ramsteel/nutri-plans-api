package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) error
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

func (u *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	user := &entities.User{AuthID: id}
	err := u.db.Preload("Auth.RoleType").Preload(clause.Associations).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	return u.db.Omit(
		"Auth.Password",
	).Session(&gorm.Session{FullSaveAssociations: true}).Updates(user).Error
}
