package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HandleKey struct {
	Path   string
	Method string
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[HandleKey]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[HandleKey]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func NewHandleKey(method, path string) *HandleKey {
	return &HandleKey{
		Path:   path,
		Method: method,
	}
}

func (s *WebServer) AddHandler(handleKey *HandleKey, handler http.HandlerFunc) {
	s.Handlers[*handleKey] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for handleKey, handler := range s.Handlers {
		switch handleKey.Method {
		case "GET":
			s.Router.Get(handleKey.Path, handler)
		case "POST":
			s.Router.Post(handleKey.Path, handler)
		case "PUT":
			s.Router.Put(handleKey.Path, handler)
		case "DELETE":
			s.Router.Delete(handleKey.Path, handler)
		case "PATCH":
			s.Router.Patch(handleKey.Path, handler)
		default:
			s.Router.Get(handleKey.Path, handler)
		}

	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
