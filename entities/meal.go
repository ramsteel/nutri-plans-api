package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Meal struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID uuid.UUID `json:"-" gorm:"type:uuid;index"`
	User   User      `json:"-"`
	CalculatedNutrients
	MealItems []MealItem     `json:"meal_items"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
