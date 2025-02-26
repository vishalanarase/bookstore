package middleware

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

var limiter = NewIPRateLimiter(1, 5)

// IPRateLimiter returns a IP address wise rate limiter
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter initiates and resturns a new IPRateLimiter
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	return i
}

// AddIP creates a new rate limiter and adds it to the ips map,
// using the IP address as the key
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = limiter

	return limiter
}

// GetLimiter returns the rate limiter for the provided IP address if it exists.
// Otherwise calls AddIP to add IP address to the map
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	if strings.Contains(ip, ":") {
		ip = strings.Split(ip, ":")[0]
	}
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}

	i.mu.Unlock()
	return limiter
}

// RateLimitHandler validates the rate limit
func RateLimitHandler(context *gin.Context) {
	lmt := limiter.GetLimiter(context.ClientIP())

	if !lmt.Allow() {
		log.Errorf("Too many requests from %s", context.Request.RemoteAddr)
		context.AbortWithStatusJSON(http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
	}
	context.Next()
}
