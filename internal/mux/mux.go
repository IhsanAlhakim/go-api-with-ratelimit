package mux

import (
	"net/http"
)

type Mux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (m *Mux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	m.middlewares = append(m.middlewares, next)
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &m.ServeMux

	/*
		middleware paling awal di slice adalah middleware yg membungkus pertama handler
		jadi ketika register middleware 1 baru 2
		yg dijalanin yg kedua dulu karna paling luar
	*/
	// for i, next := range m.middlewares {
	// 	current = next(current)
	// }

	/*
		kalau mau yg registrasi pertama dijalanin pertama
	*/
	for i := len(m.middlewares) - 1; i >= 0; i-- {
		current = m.middlewares[i](current)
	}

	current.ServeHTTP(w, r)
}

func New() *Mux {
	return &Mux{}
}
