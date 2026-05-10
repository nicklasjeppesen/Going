package db

import (
	"fmt"
	. "myapp/internal/app/models/appable"
	"strconv"

	. "github.com/nicklasjeppesen/going_internal/super/db"
	. "github.com/nicklasjeppesen/going_internal/super/db/types"
)

type Company struct {
	ActiveRecord[*Company] `json:"-" hidden:"true" swaggerignore:"true" validate:"-"`
	Name                   string                 `json:"name" validate:"required"`
	ConcernTest            ConcernsImplementation `validate:"-" json:"-"`

	// DelegateAbles
	EntryAble `validate:"-" json:"-"` // include the EntryAble struct to make company delegatable
}

func (c Company) DB() *Company {

	company := &c
	company.Table = "companies"
	company.Columns = company.columns()
	company.ParentDB = CreateORM(company)
	company.ConcernTest = ConcernsImplementation{Concern{company}}
	company.CustomScope(2)
	return company
}

func (company *Company) columns() Columns {
	return Columns{"name": &company.Name}
}

func (company *Company) Users() HasMany[*User] { return NewHasMany(User{}.DB(), company) }
func (company *Company) User() HasOne[*User]   { return NewHasOne(User{}.DB(), company) }
func (Company *Company) Comments() HasManyMorph[*Comment] {
	return NewHasManyMorph(Comment{}.DB(), Company)
}

func (company *Company) WithComments() *Company {
	return company.With("Comments")
}

// ------------- DB Scopes ------------------------ //
func (c *Company) CustomScope(id int) *Company { c.Where("id", id); return c }

// -------------- EntryAble functions --------------------//
func (company *Company) Title() string {
	return company.Name + ", Company Title is called"
}

func (company *Company) CustomID() string {
	return strconv.Itoa(int(company.ActiveRecord.Id))
}

func (company *Company) PrintHi() string {
	return "Hi from Company: " + fmt.Sprint(company.Id)
}
