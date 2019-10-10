package engine

import (
	"fmt"
	"html/template"
	kb "kanbanBoard"
	"kanbanBoard/repository"
	"net/http"
)

// Engine contains all tasks and sends a HTML document build from templates
type Engine struct {
	board kb.Board
	repo  repository.SqliteRepository
}

// New Engine constructor
func New() *Engine {
	e := &Engine{}
	e.board = kb.Board{}
	return e
}

//save to repo
//template render

// Render and send HTML document to the server
func (e *Engine) Render(w http.ResponseWriter, r *http.Request) {

	p := "../engine/templates/"
	temp := template.Must(template.ParseFiles(p+"body.html", p+"ticket.html"))

	err := temp.Execute(w, e.board)

	fmt.Println(err)
}

// SetBoard is a setter to 'update' the Board
func (e *Engine) SetBoard(t kb.Board) {
	e.board = t
}
