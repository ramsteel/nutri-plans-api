package repositories

import (
	"context"
	roleconst "nutri-plans-api/constants/role"
	"nutri-plans-api/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateAuth(ctx context.Context, auth *entities.Auth) error
	GetAuthByEmail(ctx context.Context, email string) (*entities.Auth, error)
	GetAllUsersAuths(ctx context.Context) (*[]entities.Auth, error)
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

func (a *authRepository) GetAuthByEmail(ctx context.Context, email string) (*entities.Auth, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	auth := new(entities.Auth)
	if err := a.db.Where("email = ?", email).First(auth).Error; err != nil {
		return nil, err
	}

	return auth, nil
}

func (a *authRepository) GetAllUsersAuths(ctx context.Context) (*[]entities.Auth, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	auths := new([]entities.Auth)
	err := a.db.Preload("RoleType").
		Where("role_type_id = ?", roleconst.UserRoleID).
		Find(auths).Error

	if err != nil {
		return nil, err
	}

	return auths, nil
}
