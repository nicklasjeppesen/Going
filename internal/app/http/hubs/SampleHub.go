package hubs

import (
	"fmt"
	"myapp/internal/app/helper"

	socket "github.com/nicklasjeppesen/going_internal/super/socket"
)

type SampleHub struct {
	socket.BaseHub
	logger helper.ILogger
}

// Loader resolves SampleHub's dependencies from the container. Since it
// returns nothing, socket.Router keeps using this same *SampleHub instance —
// Loader just fills in its fields.
func (h *SampleHub) Loader(logger helper.ILogger) {
	h.logger = logger
}

/*
- Register the routes for the chat hub
*/
func (sample *SampleHub) RegisterRoutes() {

	sample.On("new_message", sample.handleNewMessage)
	// chat.On("other_event", chat.handleOtherEvent)
}

func (sample *SampleHub) handleNewMessage(parameters []string, client *socket.Client) error {
	sample.logger.Log("Received new message: " + fmt.Sprintf("%v", parameters))

	client.SendMessage("new_message", "Thank you for your message")
	return nil
}

/*
- handler when user is cancle a connection
*/
func (sample *SampleHub) CancelConnection(client *socket.Client) {
	fmt.Println("User cancle the connection")
}
