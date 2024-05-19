package rate

import (
	"net/http"
	"time"

	msgconst "nutri-plans-api/constants/message"
	httputil "nutri-plans-api/utils/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func GetRateLimiterConfig() *middleware.RateLimiterConfig {
	return &middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      rate.Limit(2) / 3600,
				Burst:     12,
				ExpiresIn: 3 * time.Minute,
			},
		),
		IdentifierExtractor: func(c echo.Context) (string, error) {
			id := c.Request().Header.Get("X-Forwarded-For")
			return id, nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return httputil.HandleErrorResponse(
				c,
				http.StatusForbidden,
				msgconst.MsgForbiddenResource,
			)
		},
		DenyHandler: func(c echo.Context, identifier string, err error) error {
			return httputil.HandleErrorResponse(
				c,
				http.StatusTooManyRequests,
				msgconst.MsgTooManyRequest,
			)
		},
	}
}
