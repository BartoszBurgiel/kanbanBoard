package engine

import (
	"html/template"
	kb "kanbanBoard"
	"kanbanBoard/repository"
	"net/http"
)

// Engine contains all tasks and sends a HTML document build from templates
type Engine struct {
	tasks kb.Tasks
	repo  repository.SqliteRepository
}

// New Engine constructor
func New() *Engine {
	e := &Engine{}
	e.tasks = kb.Tasks{}
	return e
}

//save to repo
//template render

// Render and send HTML document to the server
func (e *Engine) Render(w http.ResponseWriter, r *http.Request) {

	p := "../engine/templates/"
	temp := template.Must(template.ParseFiles(p+"body.html", p+"ticket.html"))

	temp.Execute(w, e.tasks)
}

// SetTasks is a setter to 'update' the tasks
func (e *Engine) SetTasks(t kb.Tasks) {
	e.tasks = t
}
