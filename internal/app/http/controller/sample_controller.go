// main.go
package controller

import (
	"fmt"
	"myapp/internal/app/models/db"
	"strconv"

	"github.com/nicklasjeppesen/going_internal/super/channels"
)

type SampleController struct {
	ControllerBase
}

// @Tags         Sample
// @Summary Get all users
// @Produce  json
// @Router /sample [get]
// @Success 200 {array} models.User
func (c *SampleController) Get(r Request) Result {

	// Check documentation for this

	websocketChannel.SendMessageToSocket(
		channels.Socket{
			URL:     "/ws/message",
			Message: "Message from a controller",
		})

	users := db.User{}.DB(r.R.Context()).Pagination(r.R, 3)
	return Response.PrintJson(users)
}

// @Tags         Sample
// @Summary Get all users
// @Produce  json
// @Router /user/index [get]
func (c *SampleController) ShowUser(id string) Result {
	return Response.Print(id)
}

func (c *SampleController) APIONE(r Request) Result {
	return Response.Print(r.R.URL.Path)
}

// @Tags         Sample
// @Summary Get all users
// @Produce  json
// @Router /user/index [get]
func (c *SampleController) APITWO(id string, name string) Result {
	return Response.Print(name)
}

// @Summary Store a user
// @Produce  json
// @Tags         Sample
// @Success 200 {array} models.User
// @Router /users [get]
func (c *SampleController) Store(r RequestBody[db.User]) Result {
	if result := r.Validate(); result.HasError {
		return Response.WithErrors(result.Errors).Back()

	} else {
		return Response.Print("id er følgende: " + strconv.Itoa(int(result.Data.Id)))
	}
}

type CreateUserRequest struct {
	Username string `form:"name" json:"username" validate:"required"`
	Email    string `form:"email" json:"email" validate:"required,email"`
}

func (c *SampleController) Tester(r RequestBody[CreateUserRequest]) Result {
	if result := r.Validate(); result.HasError {
		fmt.Println(result.Errors)
		return Response.PrintJson(result.Errors, 400)
	} else {
		return Response.Print("Alt er fint")
	}

}
