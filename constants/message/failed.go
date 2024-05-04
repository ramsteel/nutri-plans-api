package message

const (
	// file
	MsgFailedOpenFile = "failed to open file"

	// password
	MsgFailedHashingPassword = "failed hashing password"

	// request
	MsgMismatchedDataType = "mismatched data type"
	MsgInvalidRequestData = "invalid request data"
	MsgFailedParseDate    = "failed to parse date"

	// database
	MsgFailedConnectDB = "failed connect to database"
	MsgSeedFailed      = "database seeding failed"

	// users
	MsgUserCreationFailed = "failed to create user"
	MsgCountryNotFound    = "country not found"
	MsgUserExist          = "username or email already exist"
)
