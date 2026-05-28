package auth

import (
	"fmt"
	. "myapp/internal/app/http/controller"

	models "myapp/internal/app/models/db"

	security "github.com/nicklasjeppesen/going_internal/super/security"
)

type RegisterController struct {
	ControllerBase
}

func (login *RegisterController) RegisterGet() Result {
	return View("auth.register", Params{"Title": "Register Page"})
}

/*
- Post method for register a new user
*/
func (register *RegisterController) Register(r Request) Result {
	fmt.Println("im here")

	for key, values := range r.R.Form {
		for _, value := range values {
			fmt.Printf("%s = %s\n", key, value)
		}
	}

	user := new(models.User).DB(r.R.Context())
	user.Name = r.R.FormValue("name")
	user.Age = 30
	user.Email = r.R.FormValue("email")
	user.Password = security.HashPassword(r.R.FormValue("password"))
	user.Save()

	return View("auth.register", Params{"Title": "Register Page"}) // have to be the URL.

}
