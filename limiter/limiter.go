package limiter

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 用于保存每个 IP 的限流桶
type ipRateLimiter struct {
	mu           sync.RWMutex
	buckets      map[string]*ratelimit.Bucket
	fillInterval time.Duration
	capacity     int64
}

func newIPRateLimiter(fillInterval time.Duration, capacity int64) *ipRateLimiter {
	return &ipRateLimiter{
		buckets:      make(map[string]*ratelimit.Bucket),
		fillInterval: fillInterval,
		capacity:     capacity,
	}
}

func (l *ipRateLimiter) getBucket(ip string) *ratelimit.Bucket {
	l.mu.RLock()
	bucket, exists := l.buckets[ip]
	l.mu.RUnlock()

	if !exists {
		l.mu.Lock()
		// 创建一个令牌桶：每 fillInterval 补充一次，最大容量为 capacity
		bucket = ratelimit.NewBucket(l.fillInterval, l.capacity)
		l.buckets[ip] = bucket
		l.mu.Unlock()
	}

	return bucket
}

// RateLimitMiddleware 限流中间件工厂函数
func RateLimitMiddleware(fillInterval time.Duration, capacity int64) gin.HandlerFunc {
	limiter := newIPRateLimiter(fillInterval, capacity)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		bucket := limiter.getBucket(ip)

		if bucket.TakeAvailable(1) < 1 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			return
		}

		c.Next()
	}
}
