package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/madr/backend/internal/config"
	"golang.org/x/time/rate"
)

// rateLimiter stores rate limiters per IP
type rateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
	rate     rate.Limit
	burst    int
}

var globalLimiter *rateLimiter

// initRateLimiter initializes the global rate limiter
func initRateLimiter() {
	cfg := config.AppConfig.RateLimit
	
	// Calculate requests per second
	requestsPerSecond := float64(cfg.Requests) / cfg.Window.Seconds()
	
	globalLimiter = &rateLimiter{
		limiters: make(map[string]*rate.Limiter),
		rate:     rate.Limit(requestsPerSecond),
		burst:    cfg.Requests,
	}

	// Cleanup old limiters periodically
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			globalLimiter.cleanup()
		}
	}()
}

// getLimiter returns a rate limiter for the given IP
func (rl *rateLimiter) getLimiter(ip string) *rate.Limiter {
	rl.mu.RLock()
	limiter, exists := rl.limiters[ip]
	rl.mu.RUnlock()

	if exists {
		return limiter
	}

	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Double check after acquiring write lock
	limiter, exists = rl.limiters[ip]
	if exists {
		return limiter
	}

	// Create new limiter
	limiter = rate.NewLimiter(rl.rate, rl.burst)
	rl.limiters[ip] = limiter
	return limiter
}

// cleanup removes old limiters (simple implementation, can be improved)
func (rl *rateLimiter) cleanup() {
	// In a production system, you might want to track last access time
	// and remove limiters that haven't been used for a while
	// For now, we'll keep it simple
}

// RateLimiter is a middleware that limits the rate of requests
func RateLimiter() gin.HandlerFunc {
	// Initialize if not already done
	if globalLimiter == nil {
		initRateLimiter()
	}

	return func(c *gin.Context) {
		// Skip rate limiting if disabled
		if !config.AppConfig.RateLimit.Enabled {
			c.Next()
			return
		}

		// Get client IP
		ip := c.ClientIP()

		// Get limiter for this IP
		limiter := globalLimiter.getLimiter(ip)

		// Check if request is allowed
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

