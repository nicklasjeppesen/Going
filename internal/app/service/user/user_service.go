//  Handler for Business logic

package UserService

import (
	"errors"
	"myapp/internal/app/models/db"

	security "github.com/nicklasjeppesen/going_internal/super/security"
)

type UserService struct {
	// Add any dependencies like repositories here
}

func CreateNewUser(_user *db.User) error {
	var user = _user.DB(_user.GetCtx()).Where("email", _user.Email).First()
	if user.Any() {
		return errors.New("user already exist")
	}

	_user.Password = security.HashPassword(_user.Password)
	_user.DB(_user.GetCtx()).Save()

	return nil

}
