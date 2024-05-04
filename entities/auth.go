package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Auth struct {
	ID         uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email      string         `json:"email" gorm:"type:varchar;unique"`
	Password   string         `json:"-" gorm:"type:varchar"`
	Username   string         `json:"username" gorm:"type:varchar;unique"`
	RoleTypeID uint           `json:"-" gorm:"type:uint;default:1"`
	RoleType   RoleType       `json:"role_type"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}