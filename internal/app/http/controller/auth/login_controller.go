package auth

import (
	. "myapp/internal/app/http/controller"
	"net/http"
)

type LoginController struct {
	ControllerBase
}

// get
func (login *LoginController) LoginGet() Result {
	return View(
		"auth/login",
		Params{"Title": "Going App", "Message": "Welcome to Going"})

}

// Post
func (login *LoginController) Login(r Request) {

	auth := r.Auth()
	if auth.Attempt(map[string]any{
		"email":    r.R.FormValue("email"),
		"password": r.R.FormValue("password"),
	}) {
		http.Redirect(r.W, r.R, "/protected", 302)
	} else {
		Fail.StatusUnauthorized("Invalid credentials")
	}

}

func (login *LoginController) Protected(requst Request) Result {
	return View("protected", nil) // have to be the URL.
}

func (loginController *LoginController) Logout(r Request) {
	r.Auth().Logout()
	http.Redirect(r.W, r.R, "/auth/login", 302)
}
