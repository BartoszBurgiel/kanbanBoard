package kanbanboard

// Ticket represents a ticket on the kanban board
// every attribute except priority is a string, because data retrieved from
// POST is allways string
// Priority is the exception due to the sorting
type Ticket struct {
	Title       string
	Description string
	Deadline    string
	Priority    int
	ID          string
}

// State will be passed over to the HTML template
type State struct {
	State    string
	ID       string
	Tickets  []Ticket
	Position int
	Limit    int
}

// Board is a collection of all states with their tickets
type Board struct {
	States []State
}
