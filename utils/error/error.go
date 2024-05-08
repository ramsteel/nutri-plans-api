package error

import (
	"errors"
	msgconst "nutri-plans-api/constants/message"
)

var (
	// password
	ErrFailedHashingPassword = errors.New(msgconst.MsgFailedHashingPassword)
	ErrPasswordMismatch      = errors.New(msgconst.MsgPasswordMismatch)

	// token
	ErrFailedGeneratingToken = errors.New(msgconst.MsgFailedGeneratingToken)

	// external services
	ErrExternalService = errors.New(msgconst.MsgExternalServiceError)
	ErrItemNotFound    = errors.New(msgconst.MsgItemNotFound)
)
