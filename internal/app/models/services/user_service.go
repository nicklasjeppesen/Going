package services

import (
	appable "myapp/internal/app/models/appable"
	db "myapp/internal/app/models/db"
)

type UserService struct {
	User db.User
	appable.ConcernsImplementation
}

func (userService *UserService) Create() {

	userService.User = *db.User{}.DB()
	userService.User.Get()
	userService.ConcernsImplementation = appable.ConcernsImplementation{}

}
