package db

import (
	"context"

	. "github.com/nicklasjeppesen/going_internal/super/db"
	. "github.com/nicklasjeppesen/going_internal/super/db/types"
)

type User struct {
	ActiveRecord[*User] `json:"-" swaggerignore:"true" validate:"-"`
	Name                string `form:"name" json:"name" validate:"required"`
	Age                 int64  `form:"age" json:"age" validate:"min=0,max=99"`
	Email               string `form:"emil" json:"email" validate:"required"`
	Password            string `form:"password" json:"password" validate:"required" hidden:"true"`
	SessionToken        string `json:"-" hidden:"true"`
}

func (_user User) DB(ctx context.Context) *User {
	user := &_user
	user.Table = "users"
	user.Columns = user.columns()
	user.ParentDB = CreateORM(ctx, user)
	return user
}

func (user *User) columns() Columns {
	return Columns{
		// Column		  "values"
		"name":         &user.Name,
		"email":        &user.Email,
		"password":     &user.Password,
		"sessiontoken": &user.SessionToken,
	}
}
