package controller

import (
	helper "myapp/internal/app/helper"
	"net/http"
)

type HomeController struct {
	ControllerBase

	logger helper.ILogger
	userID string
}

// Loader defines the controller's dependencies as interfaces. The router
// resolves these from the Container and calls Loader once, at route
// registration time, to build the "real" controller instance.
func (c *HomeController) Loader(logger helper.ILogger) *HomeController {
	c.logger = logger

	c.AddBeforeAction(c.setUser).
		Only("Home", "Show", "Edit", "Update", "Delete")

	c.AddBeforeAction(c.authenticateUser).
		Except("Index", "Show")

	c.AddBeforeAction(c.authorizeAdmin).
		Except("Index", "Show")

	c.AddAfterAction(c.logRequest)

	return c
}

func (c *HomeController) setUser(w http.ResponseWriter, r *http.Request) bool {
	c.userID = "12355"
	c.logger.Log("set_user: " + c.userID)
	return true
}

func (c *HomeController) authenticateUser(w http.ResponseWriter, r *http.Request) bool {
	c.logger.Log("authenticate_user")
	// if not logged in {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return false
	// }
	return true
}

func (c *HomeController) authorizeAdmin(w http.ResponseWriter, r *http.Request) bool {
	c.logger.Log("authorize_admin")
	// if !isAdmin(r) {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return false
	// }
	return true
}

func (c *HomeController) logRequest(w http.ResponseWriter, r *http.Request) {
	c.logger.Log("done: " + r.URL.Path)
}

func (c *HomeController) Home() Result {
	c.logger.Log("HomeController: Home() called")
	return View("index", Params{"Title": "Going App", "Message": "Welcome to Going"})

}
