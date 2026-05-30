package route

import (
	middleware "myapp/internal/app/http/middleware"
	webrouter "myapp/internal/routes"
	"net/http"

	internalMiddelware "github.com/nicklasjeppesen/going_internal/super/middleware"
	"github.com/nicklasjeppesen/going_internal/super/socket"

	webstdlib "github.com/nicklasjeppesen/going_internal/super/customrouter"
)

type RouteServiceProvider struct {
}

func (route *RouteServiceProvider) Map(r *http.ServeMux) {
	mapwebRoute().RegisterRoutes(r)    // register the general Web provider
	mapSampleRoute().RegisterRoutes(r) // register new workspace
}

// Define the "web" route for the application.
func mapwebRoute() *webstdlib.MyRouter {
	return webrouter.Webrouter().
		AddmiddlewareGroup(middleware.WebMiddlewareGroup()). // Example on how to add a middleware to an entire groups of routes
		Addmiddleware(internalMiddelware.CsrfMiddleware)
}

func mapSampleRoute() *webstdlib.MyRouter {
	return webrouter.Samplerouter().
		AddmiddlewareGroup(middleware.WebMiddlewareGroup()).
		Addmiddleware(internalMiddelware.JWTMiddleware).
		Addprefix("/sample")
}

func RegisterMaps(r *http.ServeMux) {
	var route = RouteServiceProvider{}
	route.Map(r)
	socketRouter := socket.NewSocketRouter()
	var socket = webrouter.Socketrouter(socketRouter)
	socket.RegisterRoutes(r)
}
