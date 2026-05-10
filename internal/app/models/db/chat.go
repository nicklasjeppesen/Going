package db

import (
	. "github.com/nicklasjeppesen/going_internal/super/db"
	. "github.com/nicklasjeppesen/going_internal/super/db/types"
)

/*
* Chat symbolize any kind of chat, direct chat, group chat and so on.
 */

type Chat struct {
	ActiveRecord[*Chat] `json:"-" swaggerignore:"true" validate:"-"`
	Name                string   `json:"name"`
	Type                ChatType `json:"type"`
}
type ChatType = string

// Create a new DB instance for Chats model
func (_chat Chat) DB() *Chat {
	chat := &_chat
	chat.Table = "chats"
	chat.Columns = chat.columns()
	chat.ParentDB = CreateORM(chat)
	return chat
}

func (chat *Chat) columns() Columns {
	return Columns{
		// Column		  "values"
		"name": &chat.Name,
		"type": &chat.Type,
	}
}

const (
	direct ChatType = "direct"
	group  ChatType = "group"
	closed ChatType = "closed"
)

func (chat *Chat) WithUsers() *Chat    { return chat.With("Users") }
func (chat *Chat) WithMessages() *Chat { return chat.With("Messages") }

func (chat *Chat) Users() BelongsToMany[*User] {
	return NewBelongsToMany(User{}.DB(), chat)
}

func (chat *Chat) Messages() HasMany[*Message] {
	return NewHasMany(Message{}.DB().WithUser(), chat)
}
