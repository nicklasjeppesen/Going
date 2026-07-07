package route

import (
	"fmt"
	middleware "myapp/internal/app/http/middleware"
	"myapp/internal/app/providers/container"
	webrouter "myapp/internal/routes"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	internalMiddelware "github.com/nicklasjeppesen/going_internal/super/middleware"
	"github.com/nicklasjeppesen/going_internal/super/socket"

	webstdlib "github.com/nicklasjeppesen/going_internal/super/customrouter"
)

// Should be removed, but here for the example
func passwordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	fmt.Println("Passwrd Strength kaldt")
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	return hasUpper && hasLower && hasNumber
}

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
	_webrouter := webstdlib.NewMyRouter().UseContainer(container.GetContainer()).
		UseValidation("password_strength", passwordStrength)

	return webrouter.Webrouter(_webrouter).
		AddmiddlewareGroup(middleware.WebMiddlewareGroup())
}

func mapSampleRoute() *webstdlib.MyRouter {
	return webrouter.Samplerouter().
		UseValidation("password_strength", passwordStrength).
		AddmiddlewareGroup(middleware.WebMiddlewareGroup()).
		Addmiddleware(internalMiddelware.JWTMiddleware).
		Addprefix("/sample")
}
