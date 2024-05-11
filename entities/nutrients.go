package entities

type CalculatedNutrients struct {
	TotalCalories     float32 `json:"total_calories" gorm:"type:numeric(10,2)" conv:"total_calories"`
	TotalCarbohydrate float32 `json:"total_carbohydrate" gorm:"type:numeric(10,2)" conv:"total_carbohydrate"`
	TotalProtein      float32 `json:"total_protein" gorm:"type:numeric(10,2)" conv:"total_protein"`
	TotalFat          float32 `json:"total_fat" gorm:"type:numeric(10,2)" conv:"total_fat"`
	TotalCholesterol  float32 `json:"total_cholesterol" gorm:"type:numeric(10,2)" conv:"total_cholesterol"`
	TotalSugars       float32 `json:"total_sugars" gorm:"type:numeric(10,2)" conv:"total_sugars"`
}
