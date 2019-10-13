package core

// TicketElement stores all needed information to
// display a ticket on the kanban board
// every attribute except priority is a string, because data retrieved from
// POST returns allways string
// Priority is the exception due to the sorting
// StatesList holds all states -> needed to display drop list
type TicketElement struct {
	Title       string
	Description string
	Deadline    string
	Priority    int
	ID          string
	StateID     string
	Destination []struct {
		DestName string
		DestID   string
	}
}

// State will be passed over to the HTML template
type State struct {
	State    string
	ID       string
	Tickets  []TicketElement
	Position int
	Limit    int
}

// Board is a collection of all states with their tickets
type Board struct {
	States     []State
	StateCount int
}
