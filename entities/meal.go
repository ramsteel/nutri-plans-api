package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Meal struct {
	ID                uuid.UUID      `json:"-" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID            uuid.UUID      `json:"-" gorm:"type:uuid"`
	User              User           `json:"-"`
	TotalCalories     float32        `json:"total_calories"`
	TotalCarbohydrate float32        `json:"total_carbohydrate"`
	TotalProtein      float32        `json:"total_protein"`
	TotalFat          float32        `json:"total_fat"`
	TotalCholesterol  float32        `json:"total_cholesterol"`
	TotalSugars       float32        `json:"total_sugars"`
	MealItems         []MealItem     `json:"meal_items"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}

type MealItem struct {
	ID           uint64         `json:"-" gorm:"primaryKey;autoIncrement"`
	MealID       uuid.UUID      `json:"-" gorm:"type:uuid"`
	Meal         Meal           `json:"-"`
	MealTypeID   uint           `json:"-" gorm:"type:uint"`
	MealType     MealType       `json:"meal_type"`
	ItemName     string         `json:"item_name"`
	Qty          float32        `json:"qty"`
	Unit         string         `json:"unit"`
	Weight       float32        `json:"weight"`
	Calories     float32        `json:"calories"`
	Carbohydrate float32        `json:"carbohydrate"`
	Protein      float32        `json:"protein"`
	Fat          float32        `json:"fat"`
	Cholesterol  float32        `json:"cholesterol"`
	Sugars       float32        `json:"sugars"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
