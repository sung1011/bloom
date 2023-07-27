package svc

import "net/http"

const Key_Server = "tk:server"

type Server interface {
	HttpHandler() http.Handler
}
