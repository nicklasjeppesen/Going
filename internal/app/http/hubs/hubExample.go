package hubs

import (
	"fmt"

	socket "github.com/nicklasjeppesen/going_internal/super/socket"
)

type ChatHub struct {
	socket.BaseHub
}

/*
- Register the routes for the chat hub
*/
func (chat *ChatHub) RegisterRoutes() {

	chat.On("new_message", chat.handleNewMessage)
	// chat.On("other_event", chat.handleOtherEvent)
}

func (chat *ChatHub) handleNewMessage(parameters []string, client *socket.Client) error {
	client.SendMessage("new_message", "Thank you for your message")
	return nil
}

/*
- handler when user is cancle a connection
*/
func (chat *ChatHub) CancleConnecetion(client *socket.Client) {
	fmt.Println("User cancle the connection")
}
