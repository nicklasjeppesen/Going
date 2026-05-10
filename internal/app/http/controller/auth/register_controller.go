package auth

import (
	. "myapp/internal/app/http/controller"
	"myapp/internal/app/models/db"
	UserServivce "myapp/internal/app/service/user"
)

type RegisterController struct {
	ControllerBase
}

func (login *RegisterController) RegisterGet() Result {
	return View("auth/register", nil)
}

/*
- Post method for register a new user
*/
func (register *RegisterController) Register(r RequestBody[db.User]) Result {
	if result := r.Validate().And(UserServivce.CreateNewUser); !result.Error {
		return View("/TODO", nil) // have to be the URL.
	} else {
		return View("/TODO", Params{"errors": result.GetErrors()})
	}
}
