package controller

import (
	"net/http"
	"time"

	"github.com/nicklasjeppesen/going_internal/super/channels"
	"github.com/nicklasjeppesen/going_internal/super/response"
	"github.com/nicklasjeppesen/going_internal/super/view/template"

	viewProvider "myapp/internal/app/providers/view"

	_request "github.com/nicklasjeppesen/going_internal/super/request"
)

type Request = _request.Requestbase
type RequestBody[T any] = _request.RequestBodybase[T]

var View = template.TemplateView{CustomViewFunctions: viewProvider.GetCustomViewFunction}.View
var Response = response.Response{}
var Fail = response.Fail{}

type Params = map[string]any

type Result = func(http.ResponseWriter, *http.Request)

type ControllerBase struct {
}

var (
	websocketChannel = channels.WebSocketMessageProvider{}
)

func (handler *ControllerBase) SetHeader(w http.ResponseWriter, key string, value string) {
	// Set a custom header
	w.Header().Set(key, value)

}

func (handler *ControllerBase) SetCookie(w http.ResponseWriter, name string, value string) {
	// Set a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	// Write response body
	w.Write([]byte("Headers and cookie set!"))
}
