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

// ChangeTicket handles a given ticket according to it's state
func (t *Tasks) ChangeTicket(state, id string) {
	//var tickets *Tickets
	//var nextState *Tickets

	switch state {
	case "ToDo":
		t.moveTicketToInProgress(id)
		break
	case "InProgress":
		t.moveTicketToDone(id)
		break
	case "Done":
		t.deleteTask(id)
		break
	}

	// search in each slices for the given ID

}

// Delcare ticket as in progress by removing it from todo slice and moving it to
// inprogress slice
func (t *Tasks) moveTicketToInProgress(id string) {
	for i, ticket := range t.ToDo {
		if ticket.ID == id {

			fmt.Println("todo")
			fmt.Println("ticket: ", ticket)

			// Delete from current state
			t.ToDo[i] = t.ToDo[len(t.ToDo)-1]
			t.ToDo = t.ToDo[:len(t.ToDo)-1]

			// Append in next state
			t.InProgress = append(t.InProgress, ticket)

			break
		}
	}
}

// Delcare ticket as done by removing it from inprogress slice and moving it to
// done slice
func (t *Tasks) moveTicketToDone(id string) {
	for i, ticket := range t.InProgress {
		if ticket.ID == id {

			fmt.Println("todo")
			fmt.Println("ticket: ", ticket)

			// Delete from current state
			t.InProgress[i] = t.InProgress[len(t.InProgress)-1]
			t.InProgress = t.InProgress[:len(t.InProgress)-1]

			// Append in next state
			t.Done = append(t.Done, ticket)

			break
		}
	}
}

// Remove one task from state "done"
func (t *Tasks) deleteTask(id string) {

	for i, ticket := range t.Done {
		if ticket.ID == id {
			t.Done[i] = t.Done[len(t.Done)-1]
			t.Done = t.Done[:len(t.Done)-1]
		}
	}
}

// RowsToTask converts given sql rows to Tasks struct
func RowsToTask(r *sql.Rows) Tasks {

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
