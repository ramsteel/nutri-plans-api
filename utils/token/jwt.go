package token

import (
	"os"
	"time"

	errutil "nutri-plans-api/utils/error"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TokenUtil interface {
	GenerateToken(uid uuid.UUID, roleID uint) (string, error)
	GetClaims(c echo.Context) *JWTClaim
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

func (*tokenUtil) GetClaims(c echo.Context) *JWTClaim {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTClaim)
	return claims
}
