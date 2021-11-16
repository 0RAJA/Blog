package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

//针对路由进行限流
/*
主要逻辑是在 Key 方法中根据 RequestURI 切割出核心路由作为键值对名称，
并在 GetBucket 和 AddBuckets 进行获取和设置 Bucket 的对应逻辑。
*/

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() *MethodLimiter {
	return &MethodLimiter{
		Limiter: &Limiter{limiterBuckets: map[string]*ratelimit.Bucket{}},
	}
}

func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

// GetBucket 根据key获取桶
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

// AddBuckets 添加桶
func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			l.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}
	return l
}
