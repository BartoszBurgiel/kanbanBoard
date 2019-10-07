package server

import (
	"fmt"
	"net/http"
	"strings"
	"webserver/server/engine"
)

// Server struct
type Server struct {
	router *Router
	engine *engine.Engine
}

// NewServer returns new server
func NewServer() (*Server, error) {
	s := &Server{
		engine: engine.New(),
	}

	return s, s.init()
}

func (s *Server) init() error {
	s.router = NewRouter()

	// daten von Db holen?
	// start  engine <- ptr  of model

	s.router.Route("/")["GET"] = http.HandlerFunc(s.handleGET)
	s.router.Route("/")["POST"] = http.HandlerFunc(s.handleUserInput)

	return nil
}

func (s *Server) handleGET(w http.ResponseWriter, r *http.Request) {
	s.engine.Render(w, r)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if !strings.HasPrefix(url, "/style.css") {
		p := Path(r.URL.Path)
		m := Method(r.Method)

		fmt.Println(p, m)

		s.router.Route(p)[m].ServeHTTP(w, r)
	}

	// 	s.handleGETAssets(w, r)
	// }
}

func (s Server) handleGETAssets(w http.ResponseWriter, r *http.Request) {
	// http.FileServer(http.Dir("./ui")).ServeHTTP(w, r)
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./ui/"))))

}
