package entities

import (
	"time"

	"gorm.io/gorm"
)

type FoodType struct {
	ID        uint           `json:"id" gorm:"type:uint;primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"type:varchar(255);unique"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
