package middleware

import "go-api-with-ratelimit/internal/ratelimiter"

type Middleware struct {
	rl *ratelimiter.RateLimiter
}

func New(rl *ratelimiter.RateLimiter) *Middleware {
	return &Middleware{rl: rl}
}
