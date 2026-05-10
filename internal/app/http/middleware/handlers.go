package middleware

import (
	"log"
	config "myapp/internal/config/cors"
	"net/http"
	"strconv"
	"time"

	stdlibmiddleware "github.com/nicklasjeppesen/going_internal/super/middleware"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

// LoggingMiddleware logs the details of each request
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		log.Printf("Request: %s %s %s %s", strconv.Itoa(wrapped.statusCode), r.Method, r.URL.Path, time.Since(start).String())

	})
}

// AuthMiddleware checks for a valid authentication token
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func MiddlewareCors(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		stdlibmiddleware.Cors(next, config.AllowedOrigins())
		next.ServeHTTP(w, r)
	})
}
