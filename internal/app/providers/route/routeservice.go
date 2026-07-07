package route

import (
	middleware "myapp/internal/app/http/middleware"
	"myapp/internal/app/providers/container"
	webrouter "myapp/internal/routes"
	"net/http"

	internalMiddelware "github.com/nicklasjeppesen/going_internal/super/middleware"
	"github.com/nicklasjeppesen/going_internal/super/socket"

	webstdlib "github.com/nicklasjeppesen/going_internal/super/customrouter"
)

func RegisterMaps(r *http.ServeMux) {
	registerHttpRoutes(r)
	registerSocketRoutes(r)
}

func registerSocketRoutes(r *http.ServeMux) {
	socketRouter := socket.NewSocketRouter()
	socketRouter.UseContainer(container.GetContainer())
	var socket = webrouter.Socketrouter(socketRouter)
	socket.RegisterRoutes(r)
}

func registerHttpRoutes(r *http.ServeMux) {
	mapwebRoute().RegisterRoutes(r)    // register the general Web provider
	mapSampleRoute().RegisterRoutes(r) // register new workspace
}

// Define the "web" route for the application.
func mapwebRoute() *webstdlib.MyRouter {
	_webrouter := webstdlib.NewMyRouter().UseContainer(container.GetContainer())

	return webrouter.Webrouter(_webrouter).
		AddmiddlewareGroup(middleware.WebMiddlewareGroup()) //. // Example on how to add a middleware to an entire groups of routes
	//Addmiddleware(internalMiddelware.CsrfMiddleware)
}

func mapSampleRoute() *webstdlib.MyRouter {
	return webrouter.Samplerouter().
		AddmiddlewareGroup(middleware.WebMiddlewareGroup()).
		Addmiddleware(internalMiddelware.JWTMiddleware).
		Addprefix("/sample")
}
