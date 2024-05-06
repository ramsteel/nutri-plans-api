package token

import (
	"errors"
	"net/http"
	"os"

	msgconst "nutri-plans-api/constants/message"
	httputil "nutri-plans-api/utils/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTClaim struct {
	UID    uuid.UUID `json:"uid"`
	RoleID uint      `json:"role_id"`
	jwt.RegisteredClaims
}

func GetJwtConfig() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTClaim)
		},
		ErrorHandler: jwtErrorHandler,
		SigningKey:   []byte(os.Getenv("JWT_KEY")),
	}
}

func jwtErrorHandler(c echo.Context, err error) error {
	code := http.StatusUnauthorized

	if errors.Is(err, echojwt.ErrJWTInvalid) {
		return httputil.HandleErrorResponse(
			c,
			code,
			msgconst.MsgInvalidToken,
		)
	}

	return httputil.HandleErrorResponse(
		c,
		code,
		msgconst.MsgUnauthorized,
	)
}
