package auth

import (
	"fmt"
	. "myapp/internal/app/http/controller"
	"myapp/internal/app/models/db"
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

	userId := requst.Auth().GetUserId()
	user := new(db.User).DB(requst.R.Context()).Where("id", userId).First()
	return View("protected", Params{"Title": "Going App", "Username": user.Name})
}

func (loginController *LoginController) Logout(r Request) Result {
	r.Auth().Logout()
	return Response.Redirect("/login")
}
