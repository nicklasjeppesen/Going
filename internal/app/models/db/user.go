package db

import (
	"fmt"
	. "myapp/internal/app/models/appable"
	"strconv"

	. "github.com/nicklasjeppesen/going_internal/super/db"
	. "github.com/nicklasjeppesen/going_internal/super/db/types"
)

type User struct {
	ActiveRecord[*User] `json:"-" swaggerignore:"true" validate:"-"`
	Name                string `json:"name" validate:"required"`
	Age                 int64  `json:"age" validate:"min=0,max=99"`
	Email               string `json:"email" validate:"required"`
	Password            string `json:"password" validate:"required" hidden:"true"`
	SessionToken        string `json:"-" hidden:"true"`
	Company_id          int64  `json:"Company_id" validate:"required"`

	// Delegatable
	EntryAble
}

// very very very important, do not change.
// This is the right method to init the DB,
// It is also used for creating new objects in a GET request, and relationship methods
func (_user User) DB() *User {
	user := &_user
	user.Table = "users"
	user.Columns = user.columns()
	user.ParentDB = CreateORM(user)
	return user
}

func (user *User) columns() Columns {
	return Columns{
		// Column		  "values"
		"age":          &user.Age,
		"name":         &user.Name,
		"email":        &user.Email,
		"password":     &user.Password,
		"sessiontoken": &user.SessionToken,
		"company_id":   &user.Company_id,
	}
}

// ------------- Relationships --------------------- //
func (user *User) Company() BelongsTo[*Company] {
	return NewBelongsTo(Company{}.DB(), user)
}

func (user *User) Chats() BelongsToMany[*Chat] {
	return NewBelongsToMany(Chat{}.DB().WithMessages(), user).
		PivotCols("created_at", "updated_at")
}

// ------------- DB Scopes ------------------------ //
func (c *User) CustomScope(id int) *User { c.Where("id", id); return c }

/*
| ----------------------------------------------//
|			   Delegates entry methods		    //
|-----------------------------------------------//
*/
func (user *User) Title() string {
	return user.Name
}

func (user *User) CustomID() string {
	return strconv.Itoa(int(user.Id))
}

func (user *User) PrintHi() string {
	return "Hi from User with ID: " + fmt.Sprint(user.Id)
}

/*
| ----------------------------------------------//
|			   Validation 		   		        //
|-----------------------------------------------//
*/
func (user *User) Validate() error {
	fmt.Println("Just here for fun :D ")
	fmt.Println(user.Name)
	return nil
}
