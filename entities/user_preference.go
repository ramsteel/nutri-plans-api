package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPreference struct {
	UserID                  uuid.UUID              `json:"-" gorm:"column:id;primaryKey;type:uuid"`
	User                    User                   `json:"-"`
	FoodTypeID              *uint                  `json:"-" gorm:"default:null" conv:"food_type_id"`
	FoodType                *FoodType              `json:"food_type"`
	DrinkTypeID             *uint                  `json:"-" gorm:"default:null" conv:"drink_type_id"`
	DrinkType               *DrinkType             `json:"drink_type"`
	DietaryPreferenceTypeID *uint                  `json:"-" gorm:"default:null" conv:"dietary_preference_type_id"`
	DietaryPreferenceType   *DietaryPreferenceType `json:"dietary_preference_type"`
	DietaryRestrictions     *[]DietaryRestriction  `json:"dietary_restrictions"`
	Recommendations         *[]Recommendation      `json:"-"`
	CreatedAt               time.Time              `json:"-"`
	UpdatedAt               time.Time              `json:"-"`
	DeletedAt               gorm.DeletedAt         `json:"-" gorm:"index"`
}
