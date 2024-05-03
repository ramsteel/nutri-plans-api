package error

import (
	"errors"
	msgconst "nutri-plans-api/constants/message"
)

var (
	ErrFailedHashingPassword = errors.New(msgconst.MsgFailedHashingPassword)
)
