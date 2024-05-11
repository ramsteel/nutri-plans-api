package message

const (
	// file
	MsgFailedOpenFile = "failed to open file"

	// password
	MsgFailedHashingPassword = "failed hashing password"
	MsgPasswordMismatch      = "password mismatch"

	// request
	MsgMismatchedDataType = "mismatched data type"
	MsgInvalidRequestData = "invalid request data"

	// database
	MsgFailedConnectDB = "failed connect to database"
	MsgSeedFailed      = "database seeding failed"
	MsgFailedMigrateDB = "failed to migrate database"

	// users
	MsgUserCreationFailed = "failed to create user"
	MsgUserExist          = "username or email already exist"
	MsgLoginFailed        = "failed to login"
	MsgUnregisteredEmail  = "unregistered email"
	MsgUnregisteredUser   = "unregistered user"
	MsgGetUserFailed      = "failed to retreive user"
	MsgUpdateUserFailed   = "failed to update user"
	MsgGetAllUsersFailed  = "failed to get all users"

	// user preference
	MsgUpdatePreferenceFailed  = "failed to update preference"
	MsgGetPreferenceFailed     = "failed to get preference"
	MsgPreferenceInputNotFound = "food/drink/dietary preference types not found"

	// token
	MsgFailedGeneratingToken = "failed generating token"
	MsgUnauthorized          = "unauthorized user"
	MsgInvalidToken          = "invalid token"

	// country
	MsgCountryNotFound    = "country not found"
	MsgGetCountriesFailed = "failed to get countries"

	// food types
	MsgGetFoodTypesFailed   = "failed to get food types"
	MsgCreateFoodTypeFailed = "failed to create food type"
	MsgFoodTypeExist        = "food type already exist"
	MsgFoodTypeNotFound     = "food type not found"
	MsgUpdateFoodTypeFailed = "failed to update food type"
	MsgDeleteFoodTypeFailed = "failed to delete food type"

	// drink types
	MsgGetDrinkTypesFailed   = "failed to get drink types"
	MsgCreateDrinkTypeFailed = "failed to create drink type"
	MsgDrinkTypeExist        = "drink type already exist"
	MsgDrinkTypeNotFound     = "drink type not found"
	MsgUpdateDrinkTypeFailed = "failed to update drink type"
	MsgDeleteDrinkTypeFailed = "failed to delete drink type"

	// dietary preference types
	MsgGetDietaryPreferenceTypesFailed = "failed to get dietary preference types"
	MsgCreateDietaryPrefTypeFailed     = "failed to create dietary preference type"
	MsgDietaryPrefTypeExist            = "dietary preference type already exist"
	MsgDietaryPrefTypeNotFound         = "dietary preference type not found"
	MsgUpdateDietaryPrefTypeFailed     = "failed to update dietary preference type"
	MsgDeleteDietaryPrefTypeFailed     = "failed to delete dietary preference type"

	// meal types
	MsgGetMealTypesFailed = "failed to get meal types"

	// meals
	MsgMealNotFound         = "meal not found"
	MsgGetTodayMealFailed   = "failed to get today meal"
	MsgAddItemToMealFailed  = "failed to add item to meal"
	MsgUpdateMealItemFailed = "failed to update meal item"
	MsgMealItemNotFound     = "meal item not found"
	MsgGetMealItemFailed    = "failed to get meal item"
	MsgDeleteMealItemFailed = "failed to delete meal item"
	MsgGetMealsFailed       = "failed to get meals"
	MsgPageNotFound         = "page not found"

	// recommendation
	MsgGetRecommendationFailed     = "failed to get recommendation"
	MsgFailedAddRecommendationCron = "failed to add recommendation cron"
	MsgFailedCreateRecommendation  = "failed to create recommendation"

	// admin
	MsgAdminNotFound         = "admin not found"
	MsgGetAdminProfileFailed = "failed to get admin profile"

	// forbidden
	MsgForbiddenResource = "forbidden resource"

	// external services
	MsgExternalServiceError   = "external service error"
	MsgQueryMinimum           = "item must at least 3 characters"
	MsgSearchItemFailed       = "failed to search item"
	MsgItemNotFound           = "item not found"
	MsgGetItemNutritionFailed = "failed to get item nutrition"
)
