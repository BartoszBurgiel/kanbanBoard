package server

import (
	"fmt"
	"kanbanBoard/engine"
	"kanbanBoard/repository"
	"net/http"
	"strings"
)

// Server struct
type Server struct {
	router *Router
	engine *engine.Engine
	repo   repository.SqliteRepository
}

// NewServer returns new server
func NewServer(r repository.SqliteRepository) (*Server, error) {
	s := &Server{
		engine: engine.New(),
		repo:   r,
	}

	return s, s.init()
}

func (s *Server) init() error {
	s.router = NewRouter()

	allTasks := s.repo.GetBoard()
	s.engine.SetBoard(allTasks)

	s.router.Route("/")["GET"] = http.HandlerFunc(s.handleIndexGET)
	s.router.Route("/")["POST"] = http.HandlerFunc(s.handleUserInput)

	s.router.Route("/manage")["GET"] = http.HandlerFunc(s.handleManageGet)

	return nil
}

// if client sends a get request -> render main page
func (s *Server) handleIndexGET(w http.ResponseWriter, r *http.Request) {
	s.engine.RenderIndex(w, r)
}

// if client sends a get request -> render manage page
func (s *Server) handleManageGet(w http.ResponseWriter, r *http.Request) {
	s.engine.RenderManage(w, r)
}

// ServeHTTP to the server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if !prefixChecker(url, "/style", "/favicon.ico") {
		p := Path(r.URL.Path)
		m := Method(r.Method)

		fmt.Println(p, m)

		s.router.Route(p)[m].ServeHTTP(w, r)
		return
	}

	s.handleGETAssets(w, r)
}

func (s Server) handleGETAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("../server/assets/")).ServeHTTP(w, r)
}

// prefixChecker checks if any of given prefixes is in the url
func prefixChecker(url string, prefix ...string) bool {
	out := false
	for _, p := range prefix {
		if strings.HasPrefix(url, p) {
			out = true
		}
	}
	return out
}
