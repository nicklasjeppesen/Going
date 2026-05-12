package auth

import (
	"fmt"
	. "myapp/internal/app/http/controller"
	requests "myapp/internal/app/http/requests"
)

type LoginController struct {
	ControllerBase
}

// get
func (login *LoginController) LoginGet() Result {
	return View("auth/login", Params{"Title": "Going App", "Message": "Welcome to Going"})

}

// Post
func (login *LoginController) Login(r RequestBody[requests.UserRequest]) Result {
	auth := r.Auth()
	failedAuth := auth.Attempt(map[string]any{
		"email":    r.Body.Email,
		"password": r.Body.Password,
	})

	if failedAuth {
		return Fail.StatusUnauthorized("Invalid credentials")
	}

	return View("/<TODO>", nil) // have to be the URL.

}

func (login *LoginController) Protected(requst Request) {

	fmt.Println("Passed all middlewares")
	requst.PrintJson("Welcome to procted area")
}

func (loginController *LoginController) Logout(r Request) Result {
	r.Auth().Logout()
	return View("auth/login", nil)
}
