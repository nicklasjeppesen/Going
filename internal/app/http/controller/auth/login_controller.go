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
	auth.Email = r.R.FormValue("email")
	auth.Password = r.R.FormValue("password")

	if auth.Attempt() {
		http.Redirect(r.W, r.R, "/protected", 302)
	} else {
		Fail.StatusUnauthorized("Invalid credentials")
		http.Redirect(r.W, r.R, "/login", 302)
	}
}

func (login *LoginController) Protected(requst Request) Result {
	return View("protected", nil) // have to be the URL.
}

func (loginController *LoginController) Logout(r Request) {
	r.Auth().Logout()
	http.Redirect(r.W, r.R, "/auth/login", 302)
}
