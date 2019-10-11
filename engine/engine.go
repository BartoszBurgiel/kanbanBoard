package engine

import (
	"fmt"
	"html/template"
	"kanbanBoard/core"
	"kanbanBoard/repository"
	"net/http"
)

// Engine contains all tasks and sends a HTML document build from templates
type Engine struct {
	board core.Board
	repo  repository.SqliteRepository
}

// New Engine constructor
func New() *Engine {
	e := &Engine{}
	e.board = core.Board{}
	return e
}

// RenderIndex and send HTML document to the server
func (e *Engine) RenderIndex(w http.ResponseWriter, r *http.Request) {

	p := "../engine/templates/index/"
	temp := template.Must(template.ParseFiles(p+"body.html", p+"ticket.html"))
	err := temp.Execute(w, e.board)
	fmt.Println(err)
}

// RenderManage and send HTML document to the server
func (e *Engine) RenderManage(w http.ResponseWriter, r *http.Request) {
	p := "../engine/templates/manage/"
	temp := template.Must(template.ParseFiles(p + "index.html"))
	err := temp.Execute(w, e.board)
	fmt.Println(err)
}

// SetBoard is a setter to 'update' the Board
func (e *Engine) SetBoard(t core.Board) {
	e.board = t
}

// GetBoard returns current board
func (e *Engine) GetBoard() *core.Board {
	return &e.board
}

// GetState returns a state with given id -> else return error
func (e *Engine) GetState(id string) (core.State, error) {
	for _, state := range e.board.States {
		if state.ID == id {
			return state, nil
		}
	}
	return core.State{}, fmt.Errorf("No state found with id %s", id)
}
