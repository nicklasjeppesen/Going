package db

import (
	. "github.com/nicklasjeppesen/going_internal/super/db"
	. "github.com/nicklasjeppesen/going_internal/super/db/types"
)

type Comment struct {
	ActiveRecord[*Comment] `json:"-" swaggerignore:"true" validate:"-"`
	Name                   string `validate:"required"`
	Text                   string
	Commentable_id         int64
	Commentable_type       string
}

func (c Comment) DB() *Comment {

	comment := &c
	comment.Table = "comments"
	comment.Columns = comment.columns()
	comment.ParentDB = CreateORM(comment)
	return comment
}

func (comment *Comment) columns() Columns {
	return Columns{
		"name":             &comment.Name,
		"text":             &comment.Text,
		"commentable_id":   &comment.Commentable_id,
		"commentable_type": &comment.Commentable_type,
	}
}
