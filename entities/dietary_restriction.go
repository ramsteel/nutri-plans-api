package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DietaryRestriction struct {
	ID               uint64         `json:"-" gorm:"type:uint;primaryKey;autoIncrement"`
	UserPreferenceID uuid.UUID      `json:"-" gorm:"index;type:uuid"`
	Name             string         `json:"name" gorm:"type:varchar(255)"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
}
