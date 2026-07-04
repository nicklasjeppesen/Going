package controller

import (
	"fmt"
	helper "myapp/internal/app/helper"
	"net/http"
)

type HomeController struct {
	ControllerBase
	logger helper.ILogger
	UserID int
}

// Loader defines the controller's dependencies as interfaces. The router
// resolves these from the Container and calls Loader once, at route
// registration time, to build the "real" controller instance.
func (c *HomeController) New(logger helper.ILogger) *HomeController {
	c.logger = logger
	return c
}

// BeforeAction runs before every action on this controller. Returning false
// aborts the request (you're responsible for writing the response yourself).
func (c *HomeController) BeforeAction(w http.ResponseWriter, r *http.Request) bool {
	c.UserID = 123 // Example: set user ID from session or token
	c.logger.Log("before action: " + r.URL.Path)
	return true
}

// AfterAction runs after the action has completed.
func (c *HomeController) AfterAction(w http.ResponseWriter, r *http.Request) {
	c.logger.Log("after action: " + r.URL.Path)
}

func (c *HomeController) Home() Result {

	fmt.Println("userId er: ", c.UserID)
	c.logger.Log("HomeController: Home() called")
	return View("index", Params{"Title": "Going App", "Message": "Welcome to Going"})

}
