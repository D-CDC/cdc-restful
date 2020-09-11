package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Throttle - 檢查 Http 請求是否超過上限
func Throttle(maxEventsPerSec int, maxBurstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(maxEventsPerSec), maxBurstSize)

	return func(context *gin.Context) {
		// 檢查是否超過數量上限
		if limiter.Allow() {
			context.Next()
			return
		}

		// 若到達上限, 則回傳 HTTP 429
		context.Error(errors.New("Limit exceeded"))
		context.AbortWithStatus(http.StatusTooManyRequests)
	}
}
