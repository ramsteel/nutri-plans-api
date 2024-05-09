package dto

type MealItemRequest struct {
	MealTypeID   uint     `json:"meal_type_id" validate:"required,min=1,max=4"`
	ItemName     string   `json:"item_name" validate:"required"`
	Qty          float32  `json:"qty" validate:"required,gt=0"`
	Unit         string   `json:"unit" validate:"required"`
	Weight       float32  `json:"weight" validate:"required,gt=0"`
	Calories     *float32 `json:"calories" validate:"required,gte=0"`
	Carbohydrate *float32 `json:"carbohydrate" validate:"required,gte=0"`
	Protein      *float32 `json:"protein" validate:"required,gte=0"`
	Fat          *float32 `json:"fat" validate:"required,gte=0"`
	Cholesterol  *float32 `json:"cholesterol" validate:"required,gte=0"`
	Sugars       *float32 `json:"sugars" validate:"required,gte=0"`
}
