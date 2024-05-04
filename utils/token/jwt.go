package token

import (
	"os"
	"time"

	errutil "nutri-plans-api/utils/error"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaim struct {
	UID    uuid.UUID `json:"uid"`
	RoleID uint      `json:"role_id"`
	jwt.RegisteredClaims
}

type TokenUtil interface {
	GenerateToken(uid uuid.UUID, roleID uint) (string, error)
}

type tokenUtil struct{}

func NewTokenUtil() *tokenUtil {
	return &tokenUtil{}
}

func (*tokenUtil) GenerateToken(uid uuid.UUID, roleID uint) (string, error) {
	claims := JWTClaim{
		UID:    uid,
		RoleID: roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := unsignedToken.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", errutil.ErrFailedHashingPassword
	}

	return signedToken, nil
}
