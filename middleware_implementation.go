package middleware

import (
	"time"

	"github.com/labstack/echo"
)

type Implementation interface {
	IPRateLimiter(timeCount time.Duration) echo.MiddlewareFunc
}
