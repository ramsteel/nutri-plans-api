package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	AuthID    uuid.UUID      `json:"-" gorm:"column:id;primaryKey;type:uuid"`
	Auth      Auth           `json:"auth"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
