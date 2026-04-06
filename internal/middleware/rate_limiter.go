package middleware

import (
	"net"
	"net/http"
)

func (m *Middleware) RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ambil ip dari request
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "failed to get ip", http.StatusInternalServerError)
			return
		}

		limiter := m.rl.GetIPLimiter(ip)

		if !limiter.Allow() {
			http.Error(w, "Too many request", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
