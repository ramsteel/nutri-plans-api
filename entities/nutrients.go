package entities

type CalculatedNutrients struct {
	TotalCalories     float32 `json:"total_calories" gorm:"type:float4"`
	TotalCarbohydrate float32 `json:"total_carbohydrate" gorm:"type:float4"`
	TotalProtein      float32 `json:"total_protein" gorm:"type:float4"`
	TotalFat          float32 `json:"total_fat" gorm:"type:float4"`
	TotalCholesterol  float32 `json:"total_cholesterol" gorm:"type:float4"`
	TotalSugars       float32 `json:"total_sugars" gorm:"type:float4"`
}
