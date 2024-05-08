package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MealItem struct {
	ID           uint64         `json:"-" gorm:"primaryKey;autoIncrement"`
	MealID       uuid.UUID      `json:"-" gorm:"type:uuid"`
	Meal         Meal           `json:"-"`
	MealTypeID   uint           `json:"-" gorm:"type:uint"`
	MealType     MealType       `json:"meal_type"`
	ItemName     string         `json:"item_name" gorm:"type:varchar(255)"`
	Qty          float32        `json:"qty" gorm:"type:float4"`
	Unit         string         `json:"unit" gorm:"type:varchar(255)"`
	Weight       float32        `json:"weight" gorm:"type:float4"`
	Calories     float32        `json:"calories" gorm:"type:float4"`
	Carbohydrate float32        `json:"carbohydrate" gorm:"type:float4"`
	Protein      float32        `json:"protein" gorm:"type:float4"`
	Fat          float32        `json:"fat" gorm:"type:float4"`
	Cholesterol  float32        `json:"cholesterol" gorm:"type:float4"`
	Sugars       float32        `json:"sugars" gorm:"type:float4"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
