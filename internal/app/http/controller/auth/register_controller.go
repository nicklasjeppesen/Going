package auth

import (
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

type RegisterUserRequest struct {
	Name     string `form:"name" json:"username" validate:"required"`
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required,min=6,password_strenght"`
}

/*
- Post method for register a new user
*/
func (register *RegisterController) Register(r RequestBody[RegisterUserRequest]) Result {
	if result := r.Validate(); result.HasError {
		return Response.WithErrors(result.Errors).Back()
	}

	user := new(models.User).DB(r.R.Context())
	user.Name = r.Body.Name
	user.Email = r.Body.Email
	user.Password = security.HashPassword(r.Body.Password)
	user.Save()

	return View("auth.register", Params{"Title": "Register Page"}) // have to be the URL.

}
