package ratelimiter

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type RateLimiter struct {
	mu      sync.Mutex
	clients map[string]*Client
	rate    rate.Limit
	burst   int
}

func (rl *RateLimiter) cleanupClientsMap() {
	for {
		time.Sleep(time.Minute)

		rl.mu.Lock()
		for ip, c := range rl.clients {
			if time.Since(c.lastSeen) > 3*time.Minute {
				delete(rl.clients, ip)
			}
		}
		rl.mu.Unlock()

	}
}

func (rl *RateLimiter) GetIPLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	client, exist := rl.clients[ip]
	if !exist {
		limiter := rate.NewLimiter(rl.rate, rl.burst)
		rl.clients[ip] = &Client{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}
	client.lastSeen = time.Now()
	return client.limiter
}

func New(r rate.Limit, b int) *RateLimiter {
	rl := &RateLimiter{
		clients: make(map[string]*Client),
		rate:    r,
		burst:   b,
	}

	go rl.cleanupClientsMap()
	return rl
}
