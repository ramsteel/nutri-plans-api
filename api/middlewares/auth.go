package middlewares

import (
	"net/http"
	msgconst "nutri-plans-api/constants/message"
	rolecosnt "nutri-plans-api/constants/role"
	httputil "nutri-plans-api/utils/http"
	tokenutil "nutri-plans-api/utils/token"

	"github.com/labstack/echo/v4"
)

func MustAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := tokenutil.NewTokenUtil()
		claims := t.GetClaims(c)

		if claims.RoleID == rolecosnt.AdminRoleID {
			return next(c)
		}

		return httputil.HandleErrorResponse(c, http.StatusForbidden, msgconst.MsgForbiddenResource)
	}
}
