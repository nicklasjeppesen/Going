package controller

import (
	"fmt"
	helper "myapp/internal/app/helper"
	"net/http"

	customrouter "github.com/nicklasjeppesen/going_internal/super/customrouter"
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

	c.AddBeforeAction(customrouter.BeforeAction{
		Name:    "set_user",
		Only:    []string{"Home", "Show", "Edit", "Update", "Delete"},
		Handler: c.setUser,
	})
	c.AddBeforeAction(customrouter.BeforeAction{
		Name:    "authenticate_user",
		Handler: c.authenticateUser,
	})
	c.AddBeforeAction(customrouter.BeforeAction{
		Name:    "authorize_admin",
		Except:  []string{"Index", "Show"},
		Handler: c.authorizeAdmin,
	})

	c.AddAfterAction(customrouter.AfterAction{
		Name:    "log_request",
		Handler: c.logRequest,
	})

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

	fmt.Println("userId er: ", c.userID)
	c.logger.Log("HomeController: Home() called")
	return View("index", Params{"Title": "Going App", "Message": "Welcome to Going"})

}
