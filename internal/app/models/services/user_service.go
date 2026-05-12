package services

import (
	db "myapp/internal/app/models/db"
)

type UserService struct {
	User db.User
}

func (userService *UserService) Create() {

	userService.User = *db.User{}.DB()
	userService.User.Get()

}
