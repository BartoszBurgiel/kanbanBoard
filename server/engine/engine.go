package engine

import (
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
	e.tasks = e.dummyData()
	return e
}

// Generate dummy data
func (e *Engine) dummyData() Tasks {
	return Tasks{
		ToDo: Tickets{
			Ticket{Title: "Ticket1", Description: "Description of the task 1", ID: "1"},
			Ticket{Title: "Ticket2", Description: "Description of the task 2", ID: "2"},
			Ticket{Title: "Ticket3", Description: "Description of the task 3", ID: "3"},
		},

		InProgress: Tickets{
			Ticket{Title: "Ticket1 in progress", Description: "Task 1 - currently in progress", ID: "4"},
			Ticket{Title: "Ticket2 in progress", Description: "Task 2 - currently in progress", ID: "5"},
			Ticket{Title: "Ticket3 in progress", Description: "Task 3 - currently in progress", ID: "6"},
		},

		Done: Tickets{
			Ticket{Title: "Ticket1 done", Description: "Task 1 - already done", ID: "7"},
			Ticket{Title: "Ticket2d done", Description: "Task 2 - already done", ID: "8"},
			Ticket{Title: "Ticket3d done", Description: "Task 3 - already done", ID: "9"},
		},
	}
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
