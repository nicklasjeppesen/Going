package db

import (
	. "myapp/internal/app/models/appable"

	. "github.com/nicklasjeppesen/going_internal/super/db"

	//. "github.com/nicklasjeppesen/going_internal/super/DB/Relationship"
	. "github.com/nicklasjeppesen/going_internal/super/db/types"
)

/*
* GOAL:
*
* Account.find(1).entries.order(created_at: :desc).limit(50)
* Bucket.Recordings.MessageBoards.first()
 */

/*
Specifications for the goals:
Goal 1: Account.find(1).entries.order(created_at: :desc).limit(50)

	An account has an id in the entry super class, and therefore a relation,
	Then we can call order and limit as usual, Then we get all the entries as usual,
	but we need to link the type and id to the entryAble type in the the Entry struct which is a struct.

Goal 2: Bucket.Recordings.MessageBoards.first()

	Bucket.Recordings.MessageBoards.first() This return Bucket.Recordings.where(entryable_type: "Message")
*/
type Entry struct {
	ActiveRecord[*Entry] `validate:"-"` //int, created_at, updated_at
	Name                 string
	EntryAble_type       string
	EntryAble_id         int64
}

const delegate = "entryable"

var (
	entryAbles = []IEntry{
		new(Company).DB(),
		new(User).DB().With("company"),
	}
)

func (e Entry) DB() *Entry {
	entry := &Entry{}
	entry.Table = "entries"
	entry.Columns = Columns{
		"name":             &entry.Name,
		delegate + "_id":   &entry.EntryAble_id,
		delegate + "_type": &entry.EntryAble_type,
	}
	entry.ParentDB = CreateORM(entry)
	return entry
}

func (e *Entry) Entryable() BelongsToMorph[IEntry] {
	return NewBelongsToMorph(entryAbles, delegate, e)
}

// ----------- Scopes ----------------------------------------//
func (e *Entry) Company() IDB[*Entry] {
	return e.WhereMorph(delegate, Company{}.DB().GetName())
}
func (e *Entry) User() IDB[*Entry] {
	return e.WhereMorph(delegate, User{}.DB().GetName())
}
