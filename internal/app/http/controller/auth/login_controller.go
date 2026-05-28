package auth

import (
	"fmt"
	. "myapp/internal/app/http/controller"
)

type LoginController struct {
	ControllerBase
}

// get
func (login *LoginController) LoginGet() Result {
	fmt.Println("Login Called")
	return View(
		"auth.login",
		Params{"Title": "Login", "Message": "Please Login!"})

}

// Post action for login
func (login *LoginController) Login(r Request) Result {

	auth := r.Auth()
	auth.Email = r.R.FormValue("email")
	auth.Password = r.R.FormValue("password")

	if auth.Attempt() {
		return Response.Redirect("protected")

	} else {
		return Response.Back(Params{"error": "Invalid credentials"})
	}
}

func (login *LoginController) Protected(requst Request) Result {
	return View("protected") // have to be the URL.
}

func (loginController *LoginController) Logout(r Request) {
	r.Auth().Logout()
	Response.Redirect("auth.login")
}
