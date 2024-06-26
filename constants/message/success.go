package message

const (
	// user
	MsgUserCreated         = "user created successfully"
	MsgLoginSuccess        = "login successfully"
	MsgRetrieveUserSuccess = "user retreived successfully"
	MsgUpdateUserSuccess   = "user updated successfully"
	MsgGetAllUsersSuccess  = "all users retreived successfully"

	// user preference
	MsgPreferenceUpdated         = "preference updated successfully"
	MsgRetrievePreferenceSuccess = "preference retreived successfully"

	// country
	MsgRetrieveCountrySuccess = "country retreived successfully"

	// food types
	MsgRetrieveFoodTypesSuccess = "food types retreived successfully"
	MsgCreateFoodTypeSuccess    = "food type created successfully"
	MsgUpdateFoodTypeSuccess    = "food type updated successfully"
	MsgDeleteFoodTypeSuccess    = "food type deleted successfully"

	// drink types
	MsgRetrieveDrinkTypesSuccess = "drink types retreived successfully"
	MsgCreateDrinkTypeSuccess    = "drink type created successfully"
	MsgUpdateDrinkTypeSuccess    = "drink type updated successfully"
	MsgDeleteDrinkTypeSuccess    = "drink type deleted successfully"

	// dietary preference types
	MsgRetrieveDietaryPreferenceTypesSuccess = "dietary preference types retreived successfully"
	MsgCreateDietaryPrefTypeSuccess          = "dietary preference type created successfully"
	MsgUpdateDietaryPrefTypeSuccess          = "dietary preference type updated successfully"
	MsgDeleteDietaryPrefTypeSuccess          = "dietary preference type deleted successfully"

	// meal types
	MsgRetrieveMealTypesSuccess = "meal types retreived successfully"

	// meals
	MsgGetTodayMealSuccess   = "today meal retreived successfully"
	MsgAddItemToMealSuccess  = "item added to meal successfully"
	MsgUpdateMealItemSuccess = "meal item updated successfully"
	MsgGetMealItemSuccess    = "meal item retreived successfully"
	MsgDeleteMealItemSuccess = "meal item deleted successfully"
	MsgGetMealsSuccess       = "meals retreived successfully"

	// recommendation
	MsgGetRecommendationSuccess = "recommendation retreived successfully"

	// admin
	MsgGetAdminProfileSuccess = "admin profile retreived successfully"

	// external service
	MsgRetrieveItemSuccess     = "item retreived successfully"
	MsgGetItemNutritionSuccess = "item nutrition retreived successfully"
)
