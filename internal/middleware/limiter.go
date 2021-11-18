package middleware

import (
	"Blog/pkg/app"
	"Blog/pkg/errcode"
	"Blog/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok { //判断此路由是否需要限流
			count := bucket.TakeAvailable(1) //占用存储桶中立即可用的令牌的数量，返回值为删除的令牌数，如果没有可用的令牌，将会返回 0
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
