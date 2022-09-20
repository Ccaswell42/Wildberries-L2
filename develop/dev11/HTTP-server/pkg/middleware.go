package pkg

import (
	"fmt"
	"net/http"
	"time"
)

// MiddlewareLogger - middleware-функция логирования запросов
func MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("[%s] %s, %s, %s, %s\n", time.Now(), r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
		next.ServeHTTP(w, r)
	}
}
