package kanbanboard

// Ticket represents a ticket on the kanban board
// every attribute is a string, because data retrieved from
// POST is allways string
type Ticket struct {
	Title       string
	Description string
	Deadline    string
	Priority    string
	ID          string
}

// Tickets will be passed over to the HTML template
type Tickets []Ticket

// Tasks is a collection of all of the ticktes (ToDo, InProgress and Done)
type Tasks struct {
	ToDo, InProgress, Done Tickets
}
