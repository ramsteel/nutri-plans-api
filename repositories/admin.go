package repositories

import (
	"context"
	"nutri-plans-api/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdminRepository interface {
	GetAdminProfile(ctx context.Context, id uuid.UUID) (*entities.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{
		db: db,
	}
}

func (a *adminRepository) GetAdminProfile(
	ctx context.Context,
	id uuid.UUID,
) (*entities.Admin, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	admin := new(entities.Admin)
	err := a.db.Preload("Auth.RoleType").
		Preload(clause.Associations).First(admin, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}
