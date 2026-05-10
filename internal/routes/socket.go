package routes

import (
	"myapp/internal/app/http/hubs"

	"github.com/nicklasjeppesen/going_internal/super/middleware"
	socket "github.com/nicklasjeppesen/going_internal/super/socket"
)

/*
|--------------------------------------------------------------------------
| Websocket routes
|--------------------------------------------------------------------------
|
| Here is where you can register websocket routes for your application.
*/
func Socketrouter(socketRouter socket.Router) socket.Router {

	// Example: Register a websocket route for a chat hub with JWT middleware
	socketRouter.MapHub("/ws/example", &hubs.ChatHub{}, middleware.JWTMiddleware)
	return socketRouter
}
