package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPreference struct {
	UserID                  uuid.UUID              `json:"-" gorm:"column:id;primaryKey;type:uuid"`
	User                    User                   `json:"-"`
	FoodTypeID              *uint                  `json:"-" gorm:"default:null"`
	FoodType                *FoodType              `json:"food_type"`
	DrinkTypeID             *uint                  `json:"-" gorm:"default:null"`
	DrinkType               *DrinkType             `json:"drink_type"`
	DietaryPreferenceTypeID *uint                  `json:"-" gorm:"default:null"`
	DietaryPreferenceType   *DietaryPreferenceType `json:"dietary_preference_type"`
	DietaryRestrictions     *[]DietaryRestriction  `json:"dietary_restrictions"`
	CreatedAt               time.Time              `json:"-"`
	UpdatedAt               time.Time              `json:"-"`
	DeletedAt               gorm.DeletedAt         `json:"-" gorm:"index"`
}
