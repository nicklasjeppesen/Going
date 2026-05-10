package db

import (
	"fmt"

	. "github.com/nicklasjeppesen/going_internal/super/db"
	. "github.com/nicklasjeppesen/going_internal/super/db/types"
)

type Message struct {
	ActiveRecord[*Message] `json:"-" swaggerignore:"true"`
	Type                   int64
	Chat_id                int64
	Sender_id              int64
	Messages               string
}

type MessageType int64

const (
	text  MessageType = 0
	file  MessageType = 1
	image MessageType = 2
)

func (message *Message) GetType() {
	if MessageType(message.Type) == text {
		fmt.Println("everything is fine")
	} else {
		fmt.Println("vired")
	}
}

func (m Message) DB() *Message {

	message := &m
	message.Table = "messages"
	message.Columns = m.columns()
	message.ParentDB = CreateORM(message)
	return message
}

func (message *Message) columns() Columns {
	return Columns{
		// Column		  "values"
		"type":     &message.Type,
		"chat_id":  &message.Chat_id,
		"user_id":  &message.Sender_id,
		"messages": &message.Messages,
	}
}

func (message *Message) WithUser() *Message {
	return message.With("User")

}

func (message *Message) User() BelongsTo[*User] {
	return NewBelongsTo(User{}.DB(), message)
}
