package middleware

import (
	"log"
	config "myapp/internal/config/cors"
	"net/http"
	"strconv"
	"time"

	stdlibmiddleware "github.com/nicklasjeppesen/going_internal/super/middleware"
	"github.com/nicklasjeppesen/going_internal/super/request"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

// LoggingMiddleware logs the details of each request
func LoggingMiddleware(next request.Handler) request.Handler {
	return func(req *request.Requestbase) {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: req.W,
			statusCode:     http.StatusOK,
		}

		req.W = wrapped
		next(req)
		log.Printf("Request: %s %s %s %s", strconv.Itoa(wrapped.statusCode), req.R.Method, req.R.URL.Path, time.Since(start).String())

	}
}

// AuthMiddleware checks for a valid authentication token
func AuthMiddleware(next request.Handler) request.Handler {
	return func(req *request.Requestbase) {

		token := req.R.Header.Get("Authorization")
		if token != "valid-token" {
			http.Error(req.W, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(req)
	}
}

func MiddlewareCors(next request.Handler) request.Handler {
	return func(req *request.Requestbase) {
		stdlibmiddleware.Cors(next, config.AllowedOrigins())

		next(req)
	}
}
