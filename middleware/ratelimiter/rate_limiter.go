package ratelimiter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimiter(rate int, capacity int64) gin.HandlerFunc {
	limiter := ratelimit.NewBucketWithRate(float64(rate), capacity)
	return func(c *gin.Context) {
		if limiter.TakeAvailable(1) < 1 {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}
