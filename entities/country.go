package entities

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID        uint           `json:"id" gorm:"type:uint;primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"type:varchar"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
