package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Implementation interface {
	IPRateLimiter(timeCount time.Duration) echo.MiddlewareFunc
}
