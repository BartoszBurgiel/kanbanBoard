package engine

import "fmt"

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

func (t *Tickets) Set() {}
func (t *Ticket) Get()  {}

func NewTicket(title, desc string) Ticket {
	return Ticket{Title: title, Description: desc, ID: "10"}
}

func ChangeTicket(state, id string, tasks *Tasks) {
	var tickets *Tickets
	var nextState *Tickets

	switch state {
	case "ToDo":
		tickets = &tasks.ToDo
		nextState = &tasks.InProgress
		break
	case "InProgress":
		tickets = &tasks.InProgress
		nextState = &tasks.Done
		break
	case "Done":
		tickets = &tasks.Done
		nextState = nil
		break
	}

	// search in each slices for the given ID
	for i, ticket := range *tickets {
		if ticket.ID == id {

			fmt.Println("todo")
			fmt.Println("ticket: ", ticket)

			// modify

			*nextState = append(*nextState, ticket)

			if nextState == nil {
				(*tickets)[i] = (*tickets)[len(*tickets)-1]
				(*tickets) = (*tickets)[:len(*tickets)-1]

			}

			break
		}
	}
}
