package nutrition

type Item struct {
	ID    string `json:"tag_id"`
	Name  string `json:"tag_name"`
	Photo Photo  `json:"photo"`
}

type Photo struct {
	Thumb string `json:"thumb"`
}

type ItemNutrition struct {
	ItemName      string    `json:"food_name"`
	ServingQty    float32   `json:"serving_qty"`
	ServingUnit   string    `json:"serving_unit"`
	ServingWeight float32   `json:"serving_weight_grams"`
	Measures      []Measure `json:"alt_measures"`
	Photo         Photo     `json:"photo"`
	Nutrient
}

type Nutrient struct {
	Calories     float32 `json:"nf_calories"`
	Fat          float32 `json:"nf_fat"`
	Cholesterol  float32 `json:"nf_cholesterol"`
	Carbohydrate float32 `json:"nf_total_carbohydrate"`
	Sugar        float32 `json:"nf_sugars"`
	Protein      float32 `json:"nf_protein"`
}

type Measure struct {
	ServingWeight float32 `json:"serving_weight"`
	Measure       string  `json:"measure"`
	Qty           float32 `json:"qty"`
}

type SearchItemResponse struct {
	Common *[]Item `json:"common"`
}

type NutritionResponse struct {
	Foods *[]ItemNutrition `json:"foods"`
}
