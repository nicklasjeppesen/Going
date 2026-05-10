package middleware

import (
	stdlibmiddleware "github.com/nicklasjeppesen/going_internal/super/middleware"
)

// Chain applies a series of middleware to a handler

func WebMiddlewareGroup() stdlibmiddleware.MiddlewareGroup {
	return stdlibmiddleware.Chain(
		MiddlewareCors,
		LoggingMiddleware,
	)
}
