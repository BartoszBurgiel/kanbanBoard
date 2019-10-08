package engine

import (
	"html/template"
	"kanbanBoard/repository/sql"
	"net/http"
)

// Engine contains all tasks and sends a HTML document build from templates
type Engine struct {
	tasks Tasks
	repo  sql.SqliteRepository
}

// New Engine constructor
func New() *Engine {
	e := &Engine{}
	e.tasks = Tasks{}
	return e
}

//save to repo
//template render

// Render and send HTML document to the server
func (e *Engine) Render(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.New("body").Parse(body))
	temp.Execute(w, e.tasks)
}

// GetTasks is a for the tasks struct
func (e *Engine) GetTasks() *Tasks {
	return &e.tasks
}

// SetTasks is a setter to 'update' the tasks
func (e *Engine) SetTasks(t Tasks) {
	e.tasks = t
}
