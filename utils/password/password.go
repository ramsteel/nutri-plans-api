package password

import (
	errutil "nutri-plans-api/utils/error"

	"golang.org/x/crypto/bcrypt"
)

type PasswordUtil interface {
	HashPassword(password string) (string, error)
}

type passwordUtil struct{}

func NewPasswordUtil() *passwordUtil {
	return &passwordUtil{}
}

func (*passwordUtil) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errutil.ErrFailedHashingPassword
	}

	return string(bytes), nil
}
