package web

import "net/http"

// Server is an interface for server
type Server interface {
	Init()
	SetEndpoints() *http.ServeMux
	SetHandlers()
	SetMiddlewares(*http.ServeMux) http.Handler
	StartServe(string, string)
	ReturnHandler() http.Handler
}
