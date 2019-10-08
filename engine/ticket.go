package engine

import (
	"database/sql"
	"fmt"
)

// Ticket represents a ticket on the kanban board
type Ticket struct {
	Title       string
	Description string
	ID          string
}

// Tickets will be passed over to the HTML template
type Tickets []Ticket

// Tasks is a collection of all of the ticktes (ToDo, InProgress and Done)
type Tasks struct {
	ToDo, InProgress, Done Tickets
}

// NewTicket returns new ticket struct
func NewTicket(title, desc string) Ticket {
	return Ticket{Title: title, Description: desc, ID: "10"}
}

// RowsToTasks converts given sql rows to Tasks struct
func RowsToTasks(r *sql.Rows) Tasks {

	tasks := Tasks{}

	var title, desc, state, id string

	for r.Next() {

		err := r.Scan(&title, &desc, &state, &id)

		if err != nil {
			fmt.Println(err)
		}

		// Distinguish between states
		switch state {
		case "todo":
			tasks.ToDo = append(tasks.ToDo, Ticket{Title: title, Description: desc, ID: id})
			break
		case "inprogress":
			tasks.InProgress = append(tasks.InProgress, Ticket{Title: title, Description: desc, ID: id})
			break
		case "done":
			tasks.Done = append(tasks.Done, Ticket{Title: title, Description: desc, ID: id})
			break
		}
	}

	return tasks
}
