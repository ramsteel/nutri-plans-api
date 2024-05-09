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

	// users
	MsgUserCreationFailed = "failed to create user"
	MsgUserExist          = "username or email already exist"
	MsgLoginFailed        = "failed to login"
	MsgUnregisteredEmail  = "unregistered email"
	MsgUnregisteredUser   = "unregistered user"
	MsgGetUserFailed      = "failed to retreive user"
	MsgUpdateUserFailed   = "failed to update user"

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
	MsgGetFoodTypesFailed = "failed to get food types"

	// drink types
	MsgGetDrinkTypesFailed = "failed to get drink types"

	// dietary preference types
	MsgGetDietaryPreferenceTypesFailed = "failed to get dietary preference types"

	// meal types
	MsgGetMealTypesFailed = "failed to get meal types"

	// meals
	MsgMealNotFound         = "meal not found"
	MsgGetTodayMealFailed   = "failed to get today meal"
	MsgAddItemToMealFailed  = "failed to add item to meal"
	MsgUpdateMealItemFailed = "failed to update meal item"
	MsgMealItemNotFound     = "meal item not found"
	MsgGetMealItemFailed    = "failed to get meal item"

	// forbidden
	MsgForbiddenResource = "forbidden resource"

	// external services
	MsgExternalServiceError   = "external service error"
	MsgQueryMinimum           = "item must at least 3 characters"
	MsgSearchItemFailed       = "failed to search item"
	MsgItemNotFound           = "item not found"
	MsgGetItemNutritionFailed = "failed to get item nutrition"
)
