package entities

import "time"

type DietaryPreferenceType struct {
	ID          uint      `json:"id" gorm:"type:uint;primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar(255);unique"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   time.Time `json:"-" gorm:"index"`
}
