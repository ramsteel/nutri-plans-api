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
	MsgCountryNotFound    = "country not found"
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
)
