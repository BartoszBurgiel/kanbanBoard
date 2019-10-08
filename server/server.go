package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"kanbanBoard/server/engine"
)

// Server struct
type Server struct {
	router *Router
	engine *engine.Engine
	repo   *sql.DB
}

// NewServer returns new server
func NewServer(r *sql.DB) (*Server, error) {
	s := &Server{
		engine: engine.New(),
		repo:   r,
	}

	return s, s.init()
}

func (s *Server) init() error {
	s.router = NewRouter()

	// Get all todos
	toDos, err := s.repo.Query("SELECT * FROM todotasks ;")

	if err != nil {
		fmt.Println(err)
	}

	// Get all inprogress
	inprogress, err := s.repo.Query("SELECT * FROM inprogresstasks ;")

	if err != nil {
		fmt.Println(err)
	}

	// Get all done
	done, err := s.repo.Query("SELECT * FROM donetasks ;")

	if err != nil {
		fmt.Println(err)
	}

	// Build tasks
	tasks := engine.Tasks{
		ToDo:       engine.QueryToTickets(toDos),
		InProgress: engine.QueryToTickets(inprogress),
		Done:       engine.QueryToTickets(done),
	}

	s.engine.SetTasks(tasks)

	s.router.Route("/")["GET"] = http.HandlerFunc(s.handleGET)
	s.router.Route("/")["POST"] = http.HandlerFunc(s.handleUserInput)

	return nil
}

// if client sends a get request -> render page
func (s *Server) handleGET(w http.ResponseWriter, r *http.Request) {
	s.engine.Render(w, r)
}

// ServeHTTP to the server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if !strings.HasPrefix(url, "/style") {
		p := Path(r.URL.Path)
		m := Method(r.Method)

		fmt.Println(p, m)

		s.router.Route(p)[m].ServeHTTP(w, r)
		return
	}

	s.handleGETAssets(w, r)
}

func (s Server) handleGETAssets(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/css")
	http.FileServer(http.Dir("../server/html/style")).ServeHTTP(w, r)
}
