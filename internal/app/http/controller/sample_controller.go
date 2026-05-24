// main.go
package controller

import (
	"myapp/internal/app/models/db"
	UserService "myapp/internal/app/service/user"
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
	if result := r.Validate().And(UserService.CreateNewUser); result.Error {
		return Fail.StatusBadRequest(result.GetErrors())
	} else {
		return Response.Print("id er følgende: " + strconv.Itoa(int(result.Data.Id)))
	}
}
