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

	// For template setup
	Temp struct {
		State, Icon string
	}
}

// Tickets will be passed over to the HTML template
type Tickets []Ticket

// Tasks is a collection of all of the ticktes (ToDo, InProgress and Done)
type Tasks struct {
	ToDo, InProgress, Done Tickets
}
