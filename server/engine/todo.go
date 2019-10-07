package engine

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
