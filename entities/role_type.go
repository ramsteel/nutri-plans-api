package entities

import (
	"time"

	"gorm.io/gorm"
)

type RoleType struct {
	ID        uint           `json:"-" gorm:"type:uint;primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"type:varchar(25);unique"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
