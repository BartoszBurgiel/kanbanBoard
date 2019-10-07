package server

import (
	"net/http"
	"strings"
	"webserver/server/engine"
)

type Server struct {
	router *Router
	engine *engine.Engine
}

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

	s.router.Route("/api/")["GET"] = http.HandlerFunc(s.handleGET)
	s.router.Route("/api/")["POST"] = http.HandlerFunc(s.HandleUserInput)

	return nil
}

func (s *Server) handleGET(w http.ResponseWriter, r *http.Request) {
	s.engine.Render(w, r)
}

func (s *Server) HandleUserInput(w http.ResponseWriter, r *http.Request) {

	//daten manipulieren
	//set/get foo bar bazz

	/*
		if r.FormValue("ticketID") != "" {
			fmt.Println("Button pressed: ", r.FormValue("ticketID"))
		} else {
			fmt.Println("New ticket:")
			fmt.Println("title: ", r.FormValue("newTitle"))
			fmt.Println("desc: ", r.FormValue("newDescription"))
		}
	*/

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if strings.HasPrefix(url, "/api/") {
		p := Path(r.URL.Path)
		m := Method(r.Method)
		s.router.Route(p)[m].ServeHTTP(w, r)
	} else {
		s.handleGETAssets(w, r)
	}
}

func (s Server) handleGETAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./ui")).ServeHTTP(w, r)
	//http.Handle("/css", http.FileServer(http.Dir("../server/html/style/")))
}
