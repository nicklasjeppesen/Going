package controller

import (
	helper "myapp/internal/app/helper"
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

func (c *HomeController) setUser(request Request) bool {

	c.userID = "12355"
	c.logger.Log("set_user: " + c.userID)
	return true
}

func (c *HomeController) authenticateUser(request Request) bool {
	c.logger.Log("authenticate_user")
	// if not logged in {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return false
	// }
	return true
}

func (c *HomeController) authorizeAdmin(request Request) bool {
	c.logger.Log("authorize_admin")
	// if !isAdmin(r) {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return false
	// }
	return true
}

func (c *HomeController) logRequest(request Request) {
	c.logger.Log("done: " + request.R.URL.Path)
}

func (c *HomeController) Home() Result {
	c.logger.Log("HomeController: Home() called")
	return View("index", Params{"Title": "Going App", "Message": "Welcome to Going"})

}
