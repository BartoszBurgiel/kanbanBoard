package templates

// Ticket represents a ticket on the kanban board
type Ticket struct {
	Title       string
	Description string
	ID          int
}

// Tickets will be passed over to the HTML template
type Tickets []Ticket

// Tasks is a collection of all of the ticktes (ToDo, InProgress and Done)
type Tasks struct {
	ToDo, InProgress, Done Tickets
}
