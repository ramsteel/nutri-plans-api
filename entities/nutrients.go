package entities

type CalculatedNutrients struct {
	TotalCalories     float32 `json:"total_calories" gorm:"type:numeric(10,2)"`
	TotalCarbohydrate float32 `json:"total_carbohydrate" gorm:"type:numeric(10,2)"`
	TotalProtein      float32 `json:"total_protein" gorm:"type:numeric(10,2)"`
	TotalFat          float32 `json:"total_fat" gorm:"type:numeric(10,2)"`
	TotalCholesterol  float32 `json:"total_cholesterol" gorm:"type:numeric(10,2)"`
	TotalSugars       float32 `json:"total_sugars" gorm:"type:numeric(10,2)"`
}
