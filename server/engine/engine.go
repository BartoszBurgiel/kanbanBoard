package engine

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

// Engine contains all tasks and sends a HTML document build from templates
type Engine struct {
	tasks Tasks
}

// New Engine constructor
func New() *Engine {
	e := &Engine{}
	e.tasks = Tasks{}
	return e
}

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

// QueryToTickets turns a result row from the database to
// Tickets struct
func QueryToTickets(r *sql.Rows) Tickets {
	out := Tickets{}

	var tempTitle, tempDesc, tempID string

	for r.Next() {

		// Fetch row
		err := r.Scan(&tempTitle, &tempDesc, &tempID)

		// Check
		if err != nil {
			fmt.Println(err)
		}

		// Add to struct
		out = append(out, Ticket{tempTitle, tempDesc, tempID})
	}

	return out
}
