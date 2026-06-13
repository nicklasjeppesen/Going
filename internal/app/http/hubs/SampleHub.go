package hubs

import (
	"fmt"

	socket "github.com/nicklasjeppesen/going_internal/super/socket"
)

type SampleHub struct {
	socket.BaseHub
}

/*
- Register the routes for the chat hub
*/
func (sample *SampleHub) RegisterRoutes() {

	sample.On("new_message", sample.handleNewMessage)
	// chat.On("other_event", chat.handleOtherEvent)
}

func (sample *SampleHub) handleNewMessage(parameters []string, client *socket.Client) error {
	client.SendMessage("new_message", "Thank you for your message")
	return nil
}

/*
- handler when user is cancle a connection
*/
func (sample *SampleHub) CancelConnection(client *socket.Client) {
	fmt.Println("User cancle the connection")
}
